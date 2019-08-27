import React, { useState, useCallback } from 'react';
import { LayoutProps } from 'aw-react-router';
import { Layout, Icon } from 'antd';

import './index.less';
import SideBar from '@/components/SideBar';
import Breadcrumbs from './Breadcrumbs';

const { Header, Content, Sider } = Layout;

const BaseLayout: React.FC<LayoutProps> = ({ routerView }) => {
  const [collapsed, setCollapsed] = useState(false);

  const onCollapse = useCallback(() => {
    setCollapsed((c) => !c);
  }, [collapsed]);

  return (
    <Layout className="layout_base">
      <Sider collapsible collapsed={collapsed} trigger={null} style={{minHeight: '100vh'}}>
        <SideBar />
      </Sider>
      <Layout>
        <Header style={{ background: '#fff', padding: 0 }}>
          <Icon className="trigger" type={collapsed ? 'menu-unfold' : 'menu-fold'} onClick={onCollapse} />
        </Header>
        <Content>
          <Breadcrumbs />
          {
            routerView()
          }
        </Content>
      </Layout>
    </Layout>
  );
};

export default BaseLayout;
