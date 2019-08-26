import React from 'react';

import routes from "@/routes";
import ErrorBoundary from '@/components/ErrorBoundary';


const App: React.FC = () => {
  return (
    <ErrorBoundary>
      {
        routes
      }
    </ErrorBoundary>
  );
}

export default App;
