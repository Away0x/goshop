import React, { useEffect } from 'react';

import { withStoreProviders } from 'aw-react-store';
import AuthStore from '@/store/auth';
import routes from "@/routes";
import ErrorBoundary from '@/components/ErrorBoundary';
import { useRequest } from './tools/hooks/async';


const App: React.FC = () => {
  const { getUserInfoAction } = AuthStore.useStore();
  const { pending, run } = useRequest(getUserInfoAction, {});

  useEffect(() => {
    run();
  }, []);

  return (
    <ErrorBoundary>
      {
        pending
          ? <div className="common_page_placeholder">loading ...</div>
          : routes
      }
    </ErrorBoundary>
  );
}

export default withStoreProviders(App, [AuthStore]);
