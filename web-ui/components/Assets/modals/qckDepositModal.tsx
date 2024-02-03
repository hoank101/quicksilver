import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  Button,
  FormControl,
  FormLabel,
  Input,
  useDisclosure,
  Spinner,
} from '@chakra-ui/react';
import { StdFee } from '@cosmjs/stargate';
import { ChainName } from '@cosmos-kit/core';
import { useChain, useManager } from '@cosmos-kit/react';
import BigNumber from 'bignumber.js';
import { assets, chains } from 'chain-registry';
import Long from 'long';
import { ibc } from 'quicksilverjs';
import { useState, useMemo, useEffect } from 'react';

import { ChooseChain } from '@/components/react/choose-chain';
import { handleSelectChainDropdown, ChainOption, ChooseChainInfo } from '@/components/types';
import { useTx } from '@/hooks';
import { useIbcBalanceQuery } from '@/hooks/useQueries';
import { getIbcInfo, shiftDigits } from '@/utils';

export function DepositModal() {
  const { isOpen, onOpen, onClose } = useDisclosure();

  const [chainName, setChainName] = useState<ChainName | undefined>('osmosis');
  const { chainRecords, getChainLogo } = useManager();
  const [amount, setAmount] = useState<string>('');
  const [isLoading, setIsLoading] = useState(false);

  const chainOptions = useMemo(() => {
    return chainRecords
      .filter((chainRecord) => chainRecord.name === 'osmosis')
      .map((chainRecord) => ({
        chainName: chainRecord?.name,
        label: chainRecord?.chain?.pretty_name,
        value: chainRecord?.name,
        icon: getChainLogo(chainRecord.name),
      }));
  }, [chainRecords, getChainLogo]);

  useEffect(() => {
    setChainName(window.localStorage.getItem('selected-chain') || 'osmosis');
  }, []);

  const onChainChange: handleSelectChainDropdown = async (selectedValue: ChainOption | null) => {
    setChainName(selectedValue?.chainName);
    if (selectedValue?.chainName) {
      window?.localStorage.setItem('selected-chain', selectedValue?.chainName);
    } else {
      window?.localStorage.removeItem('selected-chain');
    }
  };

  const chooseChain = <ChooseChain chainName={chainName} chainInfos={chainOptions as ChooseChainInfo[]} onChange={onChainChange} />;

  const fromChain = chainName;
  const toChain = 'quicksilver';

  const { transfer } = ibc.applications.transfer.v1.MessageComposer.withTypeUrl;
  const { address } = useChain(fromChain ?? '');
  const { address: qAddress } = useChain('quicksilver');
  const { balance } = useIbcBalanceQuery(fromChain ?? '', address ?? '');
  const { tx } = useTx(fromChain ?? '');

  const onSubmitClick = async () => {
    setIsLoading(true);

    const transferAmount = new BigNumber(amount).shiftedBy(6).toString();

    const mainTokens = assets.find(({ chain_name }) => chain_name === chainName);
    const fees = chains.find(({ chain_name }) => chain_name === chainName)?.fees?.fee_tokens;
    const mainDenom = mainTokens?.assets[0].base ?? '';
    const fixedMinGasPrice = fees?.find(({ denom }) => denom === mainDenom)?.average_gas_price ?? '';
    const feeAmount = shiftDigits(fixedMinGasPrice, 6);

    const fee: StdFee = {
      amount: [
        {
          denom: mainDenom,
          amount: feeAmount.toString(),
        },
      ],
      gas: '500000',
    };

    const { source_port, source_channel } = getIbcInfo(fromChain ?? '', toChain ?? '');

    const token = {
      denom: 'ibc/635CB83EF1DFE598B10A3E90485306FD0D47D34217A4BE5FD9977FA010A5367D',
      amount: transferAmount,
    };

    const stamp = Date.now();
    const timeoutInNanos = new Long((stamp + 1.2e6) * 1e6);

    const msg = transfer({
      source_port,
      source_channel,
      sender: address ?? '',
      receiver: qAddress ?? '',
      token,
      timeout_height: {
        revision_number: new Long(1000),
        revision_height: new Long(1000),
      },

      timeout_timestamp: timeoutInNanos,
    });

    await tx([msg], {
      fee,
      onSuccess: () => {
        setAmount('');
      },
    });

    setIsLoading(false);
  };

  return (
    <>
      <Button
        _active={{
          transform: 'scale(0.95)',
          color: 'complimentary.800',
        }}
        onClick={onOpen}
        _hover={{
          bgColor: 'rgba(255,128,0, 0.25)',
          color: 'complimentary.300',
        }}
        color={'white'}
        w="full"
        variant="outline"
      >
        Deposit
      </Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent bgColor="rgb(32,32,32)">
          <ModalHeader color="white">Deposit QCK Tokens</ModalHeader>
          <ModalCloseButton color={'complimentary.900'} />
          <ModalBody>
            {/* Chain Selection Dropdown */}
            <FormControl>
              <FormLabel color={'white'}>From Chain</FormLabel>
              {chooseChain}
            </FormControl>

            {/* Amount Input */}
            <FormControl mt={4}>
              <FormLabel color="white">Amount</FormLabel>
              <Input
                _active={{
                  borderColor: 'complimentary.900',
                }}
                _selected={{
                  borderColor: 'complimentary.900',
                }}
                _hover={{
                  borderColor: 'complimentary.900',
                }}
                _focus={{
                  borderColor: 'complimentary.900',
                  boxShadow: '0 0 0 3px #FF8000',
                }}
                value={amount}
                onChange={(e) => setAmount(e.target.value)}
                color={'white'}
                placeholder="Enter amount"
              />
            </FormControl>
          </ModalBody>

          <ModalFooter>
            <Button
              _active={{
                transform: 'scale(0.95)',
                color: 'complimentary.800',
              }}
              _hover={{
                bgColor: 'rgba(255,128,0, 0.25)',
                color: 'complimentary.300',
              }}
              minW="100px"
              mr={3}
              onClick={onSubmitClick}
              disabled={!amount}
            >
              {isLoading === true && <Spinner size="sm" />}
              {isLoading === false && 'Deposit'}
            </Button>
            <Button
              _active={{
                transform: 'scale(0.95)',
                color: 'complimentary.800',
              }}
              _hover={{
                bgColor: 'rgba(255,128,0, 0.25)',
                color: 'complimentary.300',
              }}
              color="white"
              variant="ghost"
              onClick={onClose}
            >
              Cancel
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
}
