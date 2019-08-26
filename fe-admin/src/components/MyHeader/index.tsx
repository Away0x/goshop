import React, { FC } from 'react';
import { Row, Col } from 'antd';

import { withDefaultProps } from '@/tools/hoc';
import { withRouter, RouteComponentProps } from 'aw-react-router';
import { Nav } from './type';
import { safeGet } from '@/tools';

import { awRouter } from '@/routes';
import logoImg from '@/assets/images/logo.png';
import MenuItem from './MenuItem';
import UserAvatar from './UserAvatar';

import './index.less';


interface DefaultProps {
  canGoRoot: boolean
  navs: Nav[]
}

const defaultProps: DefaultProps = {
  canGoRoot: true,
  navs: [],
}

function isActiveNav(nav: Nav): boolean {
  const currentRoute = awRouter.getCurrentState();
  if (!currentRoute) { return false; }
  const meta = currentRoute.meta;
  if (!meta) { return false; }

  return meta.topNavName === safeGet(nav, 'route.meta.topNavName', '');
}

const MyHeader: FC<DefaultProps & RouteComponentProps> = ({ navs, canGoRoot, history }) => {
  /** 回到首页 */
  const gotoRootPage = () => {
    const homeRoute = awRouter.find('home');
    history.push(homeRoute ? homeRoute.fullPath : '/');
  };

  return (
    <div className="my_header_component">
      <Row className="menu_inner">

        {/* logo */}
        <Col span={3}>
          {
            canGoRoot
            ? <div className="menu_logo" title="点击回到首页" onClick={gotoRootPage}>
                <img alt="logo" src={logoImg} />
                <span className="logo_title">首页</span>
              </div>
            : <div className="menu_logo" style={{cursor: 'pointer'}}>
                <img alt="logo" src={logoImg} />
              </div>
          }

        </Col>

        {/* menu */}
        <Col span={19}>
          {
            navs.map((n, i) => {
              return (
                <div className="menu_wrapper_item" key={i} style={{width: 100 / navs.length + '%'}}>
                  <MenuItem nav={n} active={isActiveNav(n)} />
                </div>
              );
            })
          }
        </Col>

        {/* avatar */}
        <Col span={2}>
          <UserAvatar />
        </Col>

      </Row>
    </div>
  );
}

export default withDefaultProps(defaultProps, withRouter(MyHeader));
