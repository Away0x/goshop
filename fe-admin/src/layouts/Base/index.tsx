import React, { memo } from 'react';

import './index.less';

import { AWRouteLayout } from 'aw-react-router';
import { withStoreProviders } from 'aw-react-store';

import BaseLayoutStore from '@/store/base-layout';

import MyHeader from '@/components/MyHeader';
import MyFooter from '@/components/MyFooter';

import Breadcrumbs from './Breadcrumbs';
import { useNavs } from './hooks';


const BaseLayout: AWRouteLayout = memo(({ routerView }) => {
  const navs = useNavs();

  return (
    <div className="base_layout">
      <div className="header_wrapper">
        <MyHeader navs={navs} />
      </div>

      <div className="breadcrumb_wrapper">
        <Breadcrumbs />
      </div>

      <div className="main">
        {
          routerView()
        }
      </div>

      <MyFooter />
    </div>
  );
});

BaseLayout.displayName = 'BaseLayout';

export default withStoreProviders(BaseLayout, [BaseLayoutStore]);
