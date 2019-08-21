import { useState, useEffect } from 'react';

import { CommonResponse } from '@/services/type';

type RequestFunc = () => Promise<CommonResponse>;

interface RequestStatus extends CommonResponse {
  loading: boolean
}

export function useRequest<T>(request: RequestFunc, dep: any[], initialState: T): RequestStatus {
  const [data, setData] = useState<CommonResponse<T>>({
    status: true,
    message: '',
    data: initialState,
  });
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const r = async () => {
      setLoading(true);
      const resp = await request();
      setData(resp);
      setLoading(false);
    }
  }, dep);

  return {
    ...data,
    loading,
  }
}
