// 用户相关 (登录登出、权限 ...)
import { useEffect } from "react";
import { useSetState, UseSetState } from "react-hanger";

import { BaseValue, createStore } from 'aw-react-store';
import { CommonResponse } from '@/services/type';

import { USER_TYPE, ROLE_STATUS } from '@/constants';
import Authority from '@/auth/authority';
import Token from "@/auth/token";

import UserServices from '@/services/user';
import TermStorage from "@/storage/term";
import RoleStorage from "@/storage/role";


interface State {
  readonly user: Readonly<AW.UserEntity | null> // 用户信息
  readonly terms: Readonly<AW.TermType[]>       // 学期列表
  readonly currentTerm: Readonly<AW.TermType | null> // 当前选择的学期
  readonly currentRole: Readonly<AW.UserRoleType | null> // 当前角色
}

const initialState: State = {
  user: null,
  terms: [],
  currentTerm: null,
  currentRole: null,
}

export interface AuthValue extends BaseValue<State> {
  // actions
  loginAction: (username: string, password: string) => Promise<CommonResponse>
  getUserInfoAction: (force?: boolean) => Promise<CommonResponse>
  logoutAction: () => void
  setCurrentTermAction: (termid: number) => boolean
  setCurrentRoleAction: (roleid: number) => boolean

  // getters
  authGetter: () => Readonly<string[]>
  termsGetter: () => Readonly<AW.TermType[]>
  isTeacherGetter: () => boolean
  rolesGetter: () => AW.UserRoleType[]
}

function setUserToState(data: AW.UserEntityAndTermInfoType, state: UseSetState<State>) {
  const user = data.userEntity || {};
  const terms = data.termVos || [];

  Authority.loadMap(user.acl || {});
  state.setState({ user, terms });
}

/** 设置默认学期 */
async function setDefaultTerm(state: UseSetState<State>, defaultTermID?: number) {
  const terms = state.state.terms || [];
  if (!terms.length) {
    TermStorage.clean();
    return;
  }

  const localTermID = TermStorage.get();
  if (localTermID) {
    const localTerm = terms.find((t) => t.termid === localTermID);
    if (localTerm) {
      setCurrentTermAction(state)(localTerm.termid);
      return;
    }
  }

  if (defaultTermID) {
    const defaultTerm = terms.find((t) => t.termid === defaultTermID);
    if (defaultTerm) {
      setCurrentTermAction(state)(defaultTerm.termid);
      return;
    }
  }

  // localstroage 和 default 没有则发送请求获取
  const term = await UserServices.getDetaultTerm();
  setCurrentTermAction(state)(term.data.termid);
}

/** 设置默认角色 */
function setDefaultRole(state: UseSetState<State>) {
  const setRole = (roleid: number): boolean => {
    const inList = roles.find((r) => r.roleid === roleid);
    if (inList) {
      setCurrentRoleAction(state)(roleid);
      return true;
    }

    return false;
  }

  if (!state.state.user) {
    return;
  }

  const roles = state.state.user.aclrole || [];
  if (!roles.length) {
    RoleStorage.clean();
    return;
  }

  const localRole = RoleStorage.get();
  if (localRole) {
    if (setRole(localRole.roleid)) {
      return;
    }
  }

  const defaultRole = roles.find((r) => r.activestatus === ROLE_STATUS.IS);
  if (defaultRole) {
    if (setRole(defaultRole.roleid)) {
      return;
    }
  }

  const firstRole = roles[0];
  if (firstRole) {
    setRole(firstRole.roleid);
  }
}

/** 设置学期 */
function setCurrentTermAction(state: UseSetState<State>) {
  return (termid: number): boolean => {
    if (state.state.currentTerm && state.state.currentTerm.termid === termid) {
      return false;
    }
    const term = state.state.terms.find((t) => t.termid === termid);

    if (!term) { return false; }
    state.setState({ currentTerm: term });
    TermStorage.set(term.termid);
    return true;
  };
}

/** 设置角色 */
function setCurrentRoleAction(state: UseSetState<State>) {
  return (roleid: number): boolean => {
    if (!state.state.user) {
      return false;
    }
    if (state.state.currentRole && state.state.currentRole.roleid === roleid) {
      return false;
    }

    const roles = state.state.user.aclrole || [];
    const role = roles.find((r) => r.roleid === roleid);
    if (!role) { return false; }
    state.setState({ currentRole: role });
    RoleStorage.set({
      roleid: role.roleid,
      roletype: role.type,
    });
    return true;
  };
}

/** 登录 */
function loginAction(state: UseSetState<State>) {
  return async (username: string, password: string): Promise<CommonResponse> => {
    // 记录登录信息的 token 统一在 services request 方法中做了记录
    const result = await UserServices.login({
      username: username.trim(),
      password: password.trim(),
    });

    if (result.status && result.data) {
      setUserToState(result.data, state);
    }

    return result;
  }
}

/** 获取用户信息 */
function getUserInfoAction(state: UseSetState<State>) {
  // force: 覆盖用户数据
  return async (force = false): Promise<CommonResponse> => {
    if (!Token.isLogin()) {
      return { status: false, message: '用户未登录', data: null };
    }

    if (!force && state.state.user) {
      return { status: true, message: '成功', data: state.state.user };
    }

    // 记录登录信息的 token 统一在 services request 方法中做了记录
    const result = await UserServices.getUserInfo();

    if (result.status && result.data) {
      setUserToState(result.data, state);
    }

    return result;
  }
}

/** 登出 */
function logoutAction({ setState }: UseSetState<State>) {
  return () => {
    setState({
      ...initialState,
    });
    Token.logout({
      triggerGlobalCallback: true,
      reloadPage: true,
    });
  }
}

function useAuth(): AuthValue {
  const state = useSetState<State>(initialState);

  useEffect(() => {
    setDefaultTerm(state);
  }, [state.state.terms]);

  useEffect(() => {
    setDefaultRole(state);
  }, [state.state.user]);

  return {
    state: state.state,

    // actions
    loginAction: loginAction(state),
    getUserInfoAction: getUserInfoAction(state),
    logoutAction: logoutAction(state),
    setCurrentTermAction: setCurrentTermAction(state),
    setCurrentRoleAction: setCurrentRoleAction(state),

    // getters
    /** 获取权限 table */
    authGetter: () => Authority.authStore || [],
    /** 获取学期数据 */
    termsGetter: () => state.state.terms || [],
    /** 是否为老师 */
    isTeacherGetter: () => state.state.user ? state.state.user.category === USER_TYPE.TEACHER : false,
    /** 角色列表 */
    rolesGetter: () => {
      const info = state.state.user;
      if (!info) { return []; }

      return info.aclrole || [];
    },
  }
}

export default createStore(useAuth);
