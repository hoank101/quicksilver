import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  Divider,
  Tooltip,
  Text,
  Button,
  Table,
  Thead,
  Spinner,
  Image,
  Tbody,
  Tr,
  Th,
  Td,
  Icon,
  Flex,
  HStack,
  TableContainer,
  Menu,
  Fade,
  MenuButton,
  MenuItem,
  MenuList,
  Box,
  keyframes,
  Center,
} from '@chakra-ui/react';
import { useChain, useChains } from '@cosmos-kit/react';
import { SkipRouter, SKIP_API_URL, RouteResponse, TxStatusResponse } from '@skip-router/core';
import { ibc } from 'interchain-query';
import { useCallback, useState } from 'react';
import { FaInfoCircle } from 'react-icons/fa';

import { useTx } from '@/hooks';
import { useFeeEstimation } from '@/hooks/useFeeEstimation';
import { useAllBalancesQuery, useSkipAssets, useSkipReccomendedRoutes, useSkipRoutesData } from '@/hooks/useQueries';
import { useSkipExecute } from '@/hooks/useSkipExecute';
import { shiftDigits } from '@/utils';

const pulsate = keyframes`
0% { transform: scale(1); opacity: 0.6; }
50% { transform: scale(1.2); opacity: 0.8; }
100% { transform: scale(1); opacity: 0.6; }
`;

interface TransactionState {
  currentChain: string;
  status: string;
  txHash: string;
  txStatus?: TxStatusResponse;
  txIndex?: number;
}

const RewardsModal = ({ address, isOpen, onClose }: { address: string; isOpen: boolean; onClose: () => void }) => {
  const { wallet } = useChain('quicksilver');

  const walletMapping: { [key: string]: any } = {
    Keplr: window.keplr?.getOfflineSignerOnlyAmino,
    Cosmostation: window.cosmostation?.providers.keplr.getOfflineSignerOnlyAmino,
    Leap: window.leap.getOfflineSignerOnlyAmino,
  };

  const offlineSigner = wallet ? walletMapping[wallet.prettyName] : undefined;
  const skipClient = new SkipRouter({
    apiURL: SKIP_API_URL,
    getCosmosSigner: offlineSigner,
    endpointOptions: {
      endpoints: {
        'quicksilver-2': { rpc: 'https://rpc.quicksilver.zone/' },
      },
    },
  });
  const { balance, refetch } = useAllBalancesQuery('quicksilver', address);
  const [isSigning, setIsSigning] = useState<boolean>(false);
  const [isBottomVisible, setIsBottomVisible] = useState(true);

  const handleScroll = useCallback((event: React.UIEvent<HTMLDivElement>) => {
    const target = event.currentTarget;
    const isBottom = target.scrollHeight - target.scrollTop <= target.clientHeight;
    setIsBottomVisible(!isBottom);
  }, []);

  const chains = useChains(['cosmoshub', 'osmosis', 'stargaze', 'juno', 'sommelier', 'regen', 'dydx', 'quicksilver']);

  const { assets: skipAssets } = useSkipAssets('quicksilver-2');
  const balanceData = balance?.balances || [];

  // maps through the balance query to get the token details and assigns them to the correct values for the skip router
  const getMappedTokenDetails = () => {
    const denomArray = skipAssets?.['quicksilver-2'] || [];

    return balanceData
      .map((balanceItem) => {
        if (balanceItem.denom.startsWith('q') || balanceItem.denom.startsWith('aq') || balanceItem.denom.startsWith('uq')) {
          return null;
        }

        const denomDetail = denomArray.find((denomItem) => denomItem.denom === balanceItem.denom);

        if (denomDetail) {
          return {
            amount: balanceItem.amount,
            denom: denomDetail.denom,
            originDenom: denomDetail.originDenom,
            originChainId: denomDetail.originChainID,
            trace: denomDetail.trace,
            logoURI: denomDetail.logoURI,
            decimals: denomDetail.decimals,
          };
        }

        return null;
      })
      .filter(Boolean);
  };

  const tokenDetails = getMappedTokenDetails();

  // maps through the token details to get the route objects for the skip router
  const osmosisRouteObjects = tokenDetails.map((token) => ({
    sourceDenom: token?.denom ?? '',
    sourceChainId: 'quicksilver-2',
    destChainId: 'osmosis-1',
  }));

  const { routes: osmosisRoutes } = useSkipReccomendedRoutes(osmosisRouteObjects);
  // maps through the token details and the route objects to get the specific token details for the skip router
  const osmosisRoutesDataObjects = tokenDetails
    .flatMap((token, index) => {
      return osmosisRoutes[index]?.flatMap((route) => {
        return route.recommendations.map((recommendation) => {
          return {
            amountIn: token?.amount ?? '0',
            sourceDenom: token?.denom ?? '',
            sourceChainId: 'quicksilver-2',
            destDenom: recommendation.asset.denom ?? '',
            destChainId: recommendation.asset.chainID ?? '',
          };
        });
      });
    })
    .filter(Boolean) as { amountIn: string; sourceDenom: string; sourceChainId: string; destDenom: string; destChainId: string }[];
  const { routesData } = useSkipRoutesData(osmosisRoutesDataObjects);

  const { executeRoute, transactionState } = useSkipExecute(skipClient);
  console.log(transactionState);
  // uses all the data gathered to create the ibc transactions for sending assets to osmosis.
  const handleExecuteRoute = async () => {
    setIsSigning(true);

    const addresses = {
      'quicksilver-2': address,
      'osmosis-1': chains.osmosis.address,
      'cosmoshub-4': chains.cosmoshub.address,
      'stargaze-1': chains.stargaze.address,
      'sommelier-3': chains.sommelier.address,
      'regen-1': chains.regen.address,
      'juno-1': chains.juno.address,
      'dydx-mainnet-1': chains.dydx.address,
    };

    // Execute each route in sequence
    for (const route of routesData) {
      try {
        await executeRoute(route ?? ({} as RouteResponse), addresses);
      } catch (error) {
        console.error('Error executing route:', error);
        setIsSigning(false);
        return;
      }
    }

    setIsSigning(false);
  };

  const { tx } = useTx('quicksilver');
  const { transfer } = ibc.applications.transfer.v1.MessageComposer.withTypeUrl;
  const { estimateFee } = useFeeEstimation('quicksilver');

  const onSubmitClick = async () => {
    setIsSigning(true);

    const messages = [];

    for (const tokenDetail of tokenDetails) {
      if (!tokenDetails) {
        setIsSigning(false);
        return;
      }

      const [_, channel] = tokenDetail?.trace.split('/') ?? '';
      const sourcePort = 'transfer';
      const sourceChannel = channel;
      const senderAddress = address ?? '';

      const ibcToken = {
        denom: tokenDetail?.denom ?? '',
        amount: tokenDetail?.amount ?? '0',
      };

      const stamp = Date.now();
      const timeoutInNanos = (stamp + 1.2e6) * 1e6;

      const chainIdToName: { [key: string]: string } = {
        'osmosis-1': 'osmosis',
        'cosmoshub-4': 'cosmoshub',
        'stargaze-1': 'stargaze',
        'sommelier-3': 'sommelier',
        'regen-1': 'regen',
        'juno-1': 'juno',
        'dydx-mainnet-1': 'dydx',
      };

      const getChainName = (chainId: string) => {
        return chainIdToName[chainId] || chainId;
      };

      const chain = chains[getChainName(tokenDetail?.originChainId ?? '') ?? ''];
      const receiverAddress = chain?.address ?? '';

      const msg = transfer({
        sourcePort,
        sourceChannel,
        sender: senderAddress,
        receiver: receiverAddress,
        token: ibcToken,
        timeoutHeight: undefined,
        //@ts-ignore
        timeoutTimestamp: timeoutInNanos,
      });

      messages.push(msg);
    }

    try {
      const fee = await estimateFee(address, messages);

      await tx(messages, {
        fee,
        onSuccess: () => {
          setIsSigning(false);
        },
      });
      setIsSigning(false);
    } catch (error) {
      setIsSigning(false);
    }
  };

  const [destination, setDestination] = useState('');

  const chainIdToChainName: { [key: string]: string } = {
    'cosmoshub-4': 'cosmoshub',
    'quicksilver-2': 'quicksilver',
    'osmosis-1': 'osmosis',
    'regen-1': 'regen',
    'sommelier-3': 'sommelier',
    'dydx-mainnet-1': 'dydx',
    'juno-1': 'juno',
  };

  const getLogoUrl: { [key: string]: string } = {
    'cosmoshub-4': chains.cosmoshub?.logoUrl ?? '',
    'quicksilver-2': chains.quicksilver?.logoUrl ?? '',
    'osmosis-1': chains.osmosis?.logoUrl ?? '',
    'regen-1': chains.regen?.logoUrl ?? '',
    'sommelier-3': chains.sommelier?.logoUrl ?? '',
    'dydx-mainnet-1': chains.dydx?.logoUrl ?? '',
    'juno-1': chains.juno?.logoUrl ?? '',
  };

  const isActiveTransaction = ['Signing', 'Successful', 'Broadcasting', 'Tracking', 'Finished', 'STATE-COMPLETED', 'Error'].includes(
    transactionState.status,
  );

  const VisualTracker = ({ transactionState, routeResponse }: { transactionState: TransactionState; routeResponse: RouteResponse }) => {
    const currentChainIndex = routeResponse?.chainIDs?.findIndex((chainID: string) => chainID === transactionState.currentChain);
    console.log('In Rewards Modal', { transactionState });
    return (
      <Flex
        direction="column"
        justifyContent="center"
        alignItems="center"
        position="fixed"
        w={'sm'}
        bottom={10}
        right={10}
        bg="rgb(26,26,26)"
        p={5}
        borderRadius="lg"
      >
        <Flex flexDir={'row'} gap={2} width="100%" alignItems="center" mb={4}>
          {(transactionState.status === 'Signing' ||
            transactionState.status === 'Broadcasting' ||
            transactionState.status === 'Tracking') && <Spinner color="white" size="sm" />}
          <Text color="white" fontSize="lg">
            {`Transaction ${transactionState.status}`}
          </Text>
        </Flex>
        {['Signing', 'Broadcasting', 'Tracking', 'Successful'].includes(transactionState.status) && (
          <Flex direction="row" alignItems="center" mb={2}>
            {routeResponse?.chainIDs?.map((chainID, index) => (
              <Flex key={index} justifyContent={'center'} direction="row" alignItems="center" position="relative">
                {/* Conditional rendering for the pulsating circle */}
                {index === currentChainIndex && (
                  <Center
                    position="absolute"
                    w="50px"
                    h="50px"
                    borderRadius="full"
                    bg={
                      transactionState.status === 'Successful'
                        ? 'rgb(13, 169, 13)'
                        : transactionState.status === 'Broadcasting'
                          ? 'rgb(1, 90, 198)'
                          : transactionState.status === 'Tracking'
                            ? 'rgb(107, 107, 107)'
                            : transactionState.status === 'Signing'
                              ? 'rgb(107, 107, 107)'
                              : 'rgb(255, 0, 0)'
                    }
                    opacity="0.6"
                    animation={`${pulsate} 2s infinite`}
                    zIndex="-1"
                  />
                )}
                <Image opacity={index === currentChainIndex ? 1 : 0.3} src={`${getLogoUrl[chainID]}`} mx={4} alt={chainID} boxSize="45px" />
              </Flex>
            ))}
          </Flex>
        )}
        {transactionState.status === 'Error' && (
          <Text color="red.500" fontSize="md">
            Transaction failed
          </Text>
        )}
      </Flex>
    );
  };
  return (
    <>
      <Modal size={'sm'} isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />

        <ModalContent bgColor="rgb(32,32,32)">
          {isActiveTransaction &&
            routesData.map((route, index) => {
              return <VisualTracker key={index} transactionState={transactionState} routeResponse={route ?? ({} as RouteResponse)} />;
            })}
          <ModalHeader color="white" fontSize="xl">
            <HStack>
              <Text>Rewards</Text>
              <Tooltip
                label="These are tokens from fees, staking rewards, and holding qAsset's. Select a direction then use the unwind button to bridge them."
                fontSize="md"
                placement="right"
              >
                <span>
                  <Icon color={'complimentary.900'} as={FaInfoCircle} w={4} h={4} />
                </span>
              </Tooltip>
            </HStack>
            <Divider mt={3} bgColor={'cyan.500'} />
          </ModalHeader>
          <ModalCloseButton color={'complimentary.900'} />
          <ModalBody>
            {tokenDetails.length === 0 && (
              <Text color="white" textAlign={'center'} fontSize="md">
                No rewards available to claim
              </Text>
            )}
            {tokenDetails.length >= 1 && (
              <>
                <Box position="relative">
                  <TableContainer className="custom-scrollbar" maxH="260px" overflowY="auto" onScroll={handleScroll} position="relative">
                    <Table variant="simple" colorScheme="whiteAlpha" size="sm">
                      <Thead position="sticky" top={0} bg="rgb(32,32,32)" zIndex={1}>
                        <Tr>
                          <Th color="complimentary.900">Token</Th>
                          <Th color="complimentary.900" isNumeric>
                            Amount
                          </Th>
                        </Tr>
                      </Thead>
                      <Tbody overflowY={'auto'} maxH="250px">
                        {tokenDetails.map((detail, index) => (
                          <Tr key={index}>
                            <Td color="white">
                              <HStack>
                                <Image w="32px" h="32px" alt={detail?.originDenom} src={detail?.logoURI} />
                                <Text fontSize={'large'}>
                                  {detail?.originDenom
                                    ? detail.originDenom.toLowerCase().startsWith('factory/')
                                      ? (() => {
                                          const lastSegment = detail.originDenom.split('/').pop() || '';
                                          return lastSegment.startsWith('u')
                                            ? lastSegment.slice(1).toUpperCase()
                                            : lastSegment.toUpperCase();
                                        })()
                                      : detail.originDenom.slice(1).toUpperCase()
                                    : ''}
                                </Text>
                              </HStack>
                            </Td>
                            <Td fontSize={'large'} color="white" isNumeric>
                              {Number(shiftDigits(detail?.amount ?? '', -Number(detail?.decimals)))
                                .toFixed(2)
                                .toString()}
                            </Td>
                          </Tr>
                        ))}
                      </Tbody>
                    </Table>
                  </TableContainer>
                  {isBottomVisible && tokenDetails.length > 6 && (
                    <Fade in={isBottomVisible}>
                      <Box
                        position="absolute"
                        bottom="0"
                        left="0"
                        right="0"
                        height="70px"
                        bgGradient="linear(to top, #1A1A1A, transparent)"
                        pointerEvents="none"
                        zIndex="1"
                      />
                    </Fade>
                  )}
                </Box>
                <Flex justifyContent={'center'}>
                  <Button
                    mt={4}
                    _active={{ transform: 'scale(0.95)', color: 'complimentary.800' }}
                    _hover={{ bgColor: 'rgba(255,128,0, 0.25)', color: 'complimentary.300' }}
                    color="white"
                    size="sm"
                    w="160px"
                    variant="outline"
                    onClick={() => (destination === 'osmosis' ? handleExecuteRoute() : onSubmitClick())}
                  >
                    {isSigning === true && <Spinner size="sm" />}
                    {isSigning === false && 'Unwind'}
                  </Button>
                </Flex>
                <Flex justifyContent={'center'} mt={4}>
                  <Menu>
                    <MenuButton
                      as={Button}
                      _active={{ transform: 'scale(0.95)', color: 'complimentary.800' }}
                      _hover={{ bgColor: 'rgba(255,128,0, 0.25)', color: 'complimentary.300' }}
                      color="white"
                      size="sm"
                      w="130px"
                      bgColor={'rgb(26,26,26)'}
                    >
                      {destination === 'parentChains' ? 'Parent Chains' : destination === 'osmosis' ? 'Osmosis' : 'Direction'}
                    </MenuButton>
                    <MenuList maxW="130px" minWidth="130px" bgColor="#1a1a1a" color="white">
                      <MenuItem
                        onClick={() => setDestination('parentChains')}
                        bgColor={'rgb(26,26,26)'}
                        _hover={{ bg: 'complimentary.400' }}
                        _focus={{ bg: '#2a2a2a' }}
                      >
                        Parent Chains
                      </MenuItem>
                      <MenuItem
                        onClick={() => setDestination('osmosis')}
                        bgColor={'rgb(26,26,26)'}
                        _hover={{ bg: 'complimentary.400' }}
                        _focus={{ bg: '#2a2a2a' }}
                      >
                        Osmosis
                      </MenuItem>
                    </MenuList>
                  </Menu>
                </Flex>
              </>
            )}
          </ModalBody>

          <ModalFooter></ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};

export default RewardsModal;
