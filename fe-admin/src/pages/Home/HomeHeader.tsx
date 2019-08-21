import React, { FC } from 'react';
import { Link } from 'aw-react-router';
import {
  Dropdown,
  Menu,
  Icon,
} from 'antd';

import AuthStore from '@/store/auth';
import AWRouter from 'aw-react-router';

import defaultAvatar from '@/assets/images/default_teacher_avatar.png';
import logoImg from '@/assets/images/home/logo.png';
import logoutImg from '@/assets/images/home/logout.png';
import termImg from '@/assets/images/home/term_icon.png';


const HomeHeader: FC = () => {
  const dateNow = new Date().toLocaleDateString().replace(/\//g, '-');
  const store = AuthStore.useStore();
  const user = store.state.user!;
  const terms = store.state.terms || [];
  const roles = store.rolesGetter();
  const settingsPath = AWRouter.instance().getPathByName('settings');

  const selectTerm = (termid: number) => {
    if (store.setCurrentTermAction(termid)) {
      window.location.href = '/';
    }
  }

  const selectRole = (roleid: number) => {
    if (store.setCurrentRoleAction(roleid)) {
      window.location.href = '/';
    }
  }

  const termsMenu = (
    <Menu>
      {
        terms.map((t) => {
          return <Menu.Item key={t.termid} onClick={() => selectTerm(t.termid)}>
            {t.termname}
          </Menu.Item>
        })
      }
    </Menu>
  );

  const rolesMenu = (
    <Menu>
      {
        roles.map((r) => {
          return <Menu.Item key={r.roleid} onClick={() => selectRole(r.roleid)}>
            {r.name}
          </Menu.Item>
        })
      }
    </Menu>
  );

  return (
    <header>
      <div className="inner clearfix">
        <div className="menu_logo fl">
          <img alt="logo" src={logoImg} />
          <span className="logo_name">HappyClass智慧教学平台</span>
        </div>
        <div className="clearfix fr">
          <div className="clearfix">
            <div className="user_info fl">
              <Link className="user_avatar" title="点击进入个人中心" to={settingsPath}><img src={user.fullPath || defaultAvatar} alt="avatar" /></Link>
              <div className="fr">
                <p className="user_name">{user.username || '退出中...'}</p>
                {/* 角色切换 */}
                <Dropdown overlay={rolesMenu}>
                  <p>{store.state.currentRole ? store.state.currentRole.name : '请选择角色'} <Icon type="down" /></p>
                </Dropdown>
              </div>
            </div>
            <div className="user_info_line"></div>
            <div className="date_and_logout">
              <div className="date_now">{dateNow}</div>
              <div className="logout" onClick={store.logoutAction}>
                <div className="logout_img fl">
                  <img alt="logout" src={logoutImg} />
                </div>
                <span>退出</span>
              </div>
            </div>
          </div>
          <div className="choose_term">
            <div className="fl">
              <div className="fl term_icon">
                <img alt="term" src={termImg} />
              </div>
              <span className="fl term_text">选择学期</span>
              <div className="fl user_info_line term_line"></div>
            </div>
            <div className="fl">
              <Dropdown overlay={termsMenu}>
                <span className="menu_item_title clearfix">
                  <span className="item_title text_overflow fl" title={store.state.currentTerm ? store.state.currentTerm.termname : '请选择学期'}>
                    {store.state.currentTerm ? store.state.currentTerm.termname : '请选择学期'}
                  </span>
                  <Icon type="down" />
                </span>
              </Dropdown>
            </div>
          </div>
        </div>

      </div>
    </header>
  );
};

export default HomeHeader;
