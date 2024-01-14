import { ArrowForwardIcon, CloseIcon, InfoIcon } from '@chakra-ui/icons';
import {
  Box,
  Flex,
  Text,
  VStack,
  Button,
  Switch,
  useDisclosure,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalBody,
  Stack,
  ModalCloseButton,
  ModalHeader,
  HStack,
  Checkbox,
  Spinner,
} from '@chakra-ui/react';
import { Key, useState } from 'react';

import { useAuthChecker, useLiquidEpochQuery, useLiquidRewardsQuery } from '@/hooks/useQueries';
import { shiftDigits } from '@/utils';
import { quicksilver } from 'quicksilverjs';
import { useTx } from '@/hooks';
import { StdFee } from '@cosmjs/amino';
import { assets } from 'chain-registry';
import { cosmos } from 'interchain-query';
import { Grant, GenericAuthorization } from 'interchain-query/cosmos/authz/v1beta1/authz';
import { MsgGrant } from 'interchain-query/cosmos/authz/v1beta1/tx';

interface RewardsClaimInterface {
  address: string;
  onClose: () => void;
}

function transformProofs(proofs: any[]) {
  return proofs.map((proof) => ({
    key: proof.key,
    data: proof.data,
    proofOps: proof.proof_ops
      ? {
          //@ts-ignore
          ops: proof.proof_ops.ops.map((op) => ({
            type: op.type,
            key: op.key,
            data: op.data,
          })),
        }
      : undefined,
    height: proof.height,
    proofType: proof.proof_type,
  }));
}

export const RewardsClaim: React.FC<RewardsClaimInterface> = ({ address, onClose }) => {
  const { tx } = useTx('quicksilver' ?? '');

  const [transactionStatus, setTransactionStatus] = useState('Pending');
  const [isSigning, setIsSigning] = useState<boolean>(false);
  const [isError, setIsError] = useState<boolean>(false);
  const { liquidRewards, isLoading } = useLiquidRewardsQuery(address);
  const { liquidEpoch, isLoading: isEpochLoading } = useLiquidEpochQuery(address);

  const { submitClaim } = quicksilver.participationrewards.v1.MessageComposer.withTypeUrl;

  const { grant } = cosmos.authz.v1beta1.MessageComposer.withTypeUrl;

  const msgGrant = grant({
    granter: 'quick1c4vz0535677xpdksxh5um7zqqwfsw7245ppdaj',
    grantee: address,
    grant: Grant.fromPartial({
      authorization: GenericAuthorization.fromPartial({
        msg: '/quicksilver.participationrewards.v1.MsgSubmitClaim',
      }),
    }),
  });

  const mainTokens = assets.find(({ chain_name }) => chain_name === 'quicksilver');
  const mainDenom = mainTokens?.assets[0].base ?? 'uqck';

  const fee: StdFee = {
    amount: [
      {
        denom: mainDenom,
        amount: '5000',
      },
    ],
    gas: '500000',
  };

  const handleAutoClaimRewards = async (event: React.MouseEvent) => {
    event.preventDefault();
    setIsSigning(true);
    setTransactionStatus('Pending');
    try {
      const result = await tx([msgGrant], {
        fee,
        onSuccess: () => {},
      });
    } catch (error) {
      console.error('Transaction failed', error);
      setTransactionStatus('Failed');
      setIsError(true);
    } finally {
      setIsSigning(false);
    }
  };

  const handleClaimRewards = async (event: React.MouseEvent) => {
    event.preventDefault();
    setIsSigning(true);
    setTransactionStatus('Pending');

    if (!liquidEpoch || liquidEpoch.messages.length === 0) {
      console.error('No epoch data available or no messages to claim');
      setIsSigning(false);
      return;
    }

    try {
      const msgSubmitClaims = liquidEpoch.messages.map((message) => {
        const transformedProofs = transformProofs(message.proofs);
        return submitClaim({
          userAddress: message.user_address,
          zone: message.zone,
          srcZone: message.src_zone,
          claimType: message.claim_type,
          proofs: transformedProofs,
        });
      });

      const result = await tx(msgSubmitClaims, {
        fee,
        onSuccess: () => {},
      });
    } catch (error) {
      console.error('Transaction failed', error);
      setTransactionStatus('Failed');
      setIsError(true);
    } finally {
      setIsSigning(false);
    }
  };

  const [isOpen, setIsOpen] = useState(false);

  const [check, setCheck] = useState(false);

  const [autoClaimEnabled, setAutoClaimEnabled] = useState(false);

  const handleAutoClaimToggle = () => {
    setAutoClaimEnabled(!autoClaimEnabled);
  };

  const transactionHandler = autoClaimEnabled ? handleAutoClaimRewards : handleClaimRewards;

  return (
    <Box bgColor="rgb(32,32,32)" maxW={'sm'} p="4" borderRadius="lg" mb="4">
      <Flex direction="column" alignItems="flex-end">
        <CloseIcon color="white" cursor="pointer" onClick={onClose} _hover={{ color: 'complimentary.900' }} />
        <VStack alignItems="flex-start" spacing="2">
          <Text fontSize="xl" fontWeight="bold" color="white">
            Participation Rewards
          </Text>
          <Text pb={2} color="white" fontSize="md">
            Claim your participation rewards. Rewards will be sent to your wallet at the next epoch.
          </Text>
          <HStack gap={8} justifyContent={'space-between'}>
            <Checkbox
              _selected={{ bgColor: 'transparent' }}
              _active={{
                borderColor: 'complimentary.900',
              }}
              _hover={{
                borderColor: 'complimentary.900',
              }}
              _focus={{
                borderColor: 'complimentary.900',
                boxShadow: '0 0 0 3px #FF8000',
              }}
              isChecked={autoClaimEnabled}
              onChange={handleAutoClaimToggle}
              colorScheme="orange"
            >
              <Text color="white" fontSize="sm">
                Enable Automatic Claiming
              </Text>
            </Checkbox>
            <Button
              _hover={{
                bgColor: 'complimentary.500',
              }}
              minW={'120px'}
              onClick={transactionHandler}
              size="sm"
              alignSelf="end"
            >
              {isError ? 'Try Again' : isSigning ? <Spinner /> : 'Claim Rewards'}
            </Button>
          </HStack>
        </VStack>
      </Flex>
    </Box>
  );
};

export default RewardsClaim;
