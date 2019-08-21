// 错误边界
import React, { ErrorInfo, ComponentType } from 'react';

interface Props {
  ErrorComponent?: ComponentType
}

interface States {
  hasError: boolean
}

class ErrorBoundary extends React.Component<Props, States> {

  state = {
    hasError: false,
  }

  constructor(props: Props) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error: any) {
    console.error("ErrorBoundary: ", error);
    return { hasError: true };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    console.error("ErrorBoundary: ", {
      error, errorInfo
    });
    this.setState({
      hasError: true,
    })
  }

  render() {
    const { children, ErrorComponent } = this.props;

    if (this.state.hasError) {
      if (ErrorComponent) {
        return <ErrorComponent />
      }
      return <h1>Something went wrong.</h1>;
    }
    

    return children;
  }
}

export default ErrorBoundary;