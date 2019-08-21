import React, { FC } from 'react';
import { Link } from 'aw-react-router';
import {
  Dropdown,
  Menu,
  Icon,
} from 'antd';

import './MenuItem.less';
import { Nav } from './type';

import './UserAvatar.less';

interface Props {
  active: boolean
  nav: Nav
}

const menus = (nav: Nav): JSX.Element => {
  const subNavs = nav.subNav || [];

  return (
    <Menu>
      {
        subNavs.map((n, i) => {
          return (
            <Menu.Item key={i}>
              <Link to={n.route!.fullPath}>{n.route!.meta.hanName}</Link>
            </Menu.Item>
          );
        })
      }
    </Menu>
  );
}

const MenuItem: FC<Props> = ({ nav, active }) => {
  const classnames = [
    'menu_item_component',
    active ? 'active' : '',
  ].filter(Boolean).join(' ');
  const hasSubNav = nav.subNav && nav.subNav.length > 0 && nav.defaultName;

  return (
    <div className={classnames}>
      {
        hasSubNav
          ? <Dropdown overlayClassName="header_menu_item_content" overlay={menus(nav)}>
              <div>{nav.defaultName} <Icon type="down" /></div>
            </Dropdown>
          : <Link to={nav.route!.fullPath}>{nav.route!.meta.hanName}</Link>
      }
    </div>
  );
};

export default MenuItem;
