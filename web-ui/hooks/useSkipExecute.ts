import {SkipRouter, TxStatusResponse, RouteResponse} from '@skip-router/core';
import { useCallback, useState } from 'react';

interface TransactionState {
    currentChain: string;
    status: string;
    txHash: string;
    txStatus?: TxStatusResponse;
    txIndex?: number;
}

export function useSkipExecute(skipClient: SkipRouter) {
    if (!skipClient) {
      throw new Error('SkipRouter is not initialized');
    }
  
    const [transactionState, setTransactionState] = useState<TransactionState>({ currentChain: 'no-chain', status: 'Idle', txHash: '', txStatus: {} as TxStatusResponse, txIndex: 0});
  
    const executeRoute = useCallback(async (route: RouteResponse, userAddresses: any) => {
      setTransactionState({ currentChain: route.sourceAssetChainID, status: 'Signing', txHash: ''});
      try {
  
        await skipClient.executeRoute({
            route,
            userAddresses,
            onTransactionCompleted: async (chainID: string, txHash: string, status: TxStatusResponse) => {
                let currentChain = chainID;
                const transferSequence = status.transferSequence;
                transferSequence?.map((transferEvent) => {
                  if ('ibcTransfer' in transferEvent) {
                    currentChain = transferEvent.ibcTransfer.dstChainID;
                  }
                });
                setTransactionState({ currentChain: currentChain , status: 'Successful', txHash, txStatus: status});
            },
            onTransactionBroadcast: async (txInfo: { txHash: string; chainID: string }) => {
                setTransactionState({ currentChain: txInfo.chainID, status: 'Broadcasting', txHash: txInfo.txHash});
            },
            onTransactionTracked: async (txInfo) => {
                setTransactionState({ currentChain: txInfo.chainID, status: 'Tracking', txHash: txInfo.txHash});
            },
        });
      } catch (error) {
        console.error('Error executing route:', error);
        setTransactionState({ currentChain: 'no-chain', status: 'Error', txHash: ''});
      }
    }, [skipClient]);

    return { executeRoute, transactionState };
}

export function useSkipMessages(skipClient: SkipRouter) {
    if (!skipClient) {
        throw new Error('SkipRouter is not initialized');
    }
  const skipMessages = useCallback(async (route: any) => {
    return await skipClient.messages({
        sourceAssetDenom: route.sourceAssetDenom,
        sourceAssetChainID: route.sourceAssetChainID,
        destAssetDenom: route.destAssetDenom,
        destAssetChainID: route.destAssetChainID,
        amountIn: route.amountIn,
        amountOut: route.amountOut,
        addressList: route.addressList,
        operations: route.operations,
    });
  }, []);

  return skipMessages;
}