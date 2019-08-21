import React, { FC } from 'react';
import {
  Dropdown,
  Menu,
} from 'antd';
import { Link } from 'aw-react-router';

import defaultAvatar from '@/assets/images/default_teacher_avatar.png';
import AuthStore from '@/store/auth';
import AWRouter from 'aw-react-router';

const UserAvatar: FC = () => {
  const store = AuthStore.useStore();
  const user = store.state.user!;
  const settingsPath = AWRouter.instance().getPathByName('settings');

  const menus = (
    <Menu>
      <Menu.Item>
        <Link to={settingsPath}>个人中心</Link>
      </Menu.Item>
      <Menu.Item>
        <p onClick={store.logoutAction}>退出</p>
      </Menu.Item>
    </Menu>
  );

  return (
    <Dropdown overlay={menus}>
      <div className="user_avatar_component">
        <div className="user_avatar">
          <img src={user.fullPath || defaultAvatar} alt="avatar" />
        </div>
      </div>
    </Dropdown>
  );
};

export default UserAvatar;
