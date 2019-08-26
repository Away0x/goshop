// 权限相关中间件
import React from 'react';
import { Redirect, AWRouteState, AWMiddlewareFunc } from 'aw-react-router';

import Token from '@/auth/token';
import Authority from '@/auth/authority';

/**
 * 判断是否登录
 * whiteList: 不验证是否登录的路由 name 的白名单列表
 */
export function loginRequiredMiddleware(whiteList?: string[]): AWMiddlewareFunc {
  return function (state: AWRouteState) {
    // 路由名在白名单中，不需要验证是否登录
    if (whiteList && whiteList.length) {
      if (whiteList.indexOf(state.name) !== -1) {
        return null;
      }
    }

    if (Token.isLogin()) {
      return null;
    }

    return <Redirect to="/login" />;
  }
}

/**
 * 未登录才可访问
 * checkList: 该列表中的 route name 的路由未登录才可访问
 */
export function guestMiddleware(checkList?: string[]): AWMiddlewareFunc {
  return function (state: AWRouteState) {
    if (checkList && checkList.length) {
      if (checkList.indexOf(state.name) === -1) {
        return null;
      }
    }

    if (!Token.isLogin()) {
      return null;
    }

    return <Redirect to="/" />;
  }
}


/** 用于权限判断的中间件 (也会判断是否登录了) */
export function authCheckMiddleware() {
  return function (state: AWRouteState<AW.RouteMeta>) {
    const meta = state.meta || {};
    const auths = meta.auth || [];

    // 权限检查 (有配置才需检查)
    if (!auths.length) {
      return null;
    }
    if (Authority.checkAll(auths, !meta.authUnstrict)) {
      return null;
    }

    return <Redirect to="/unauthorized" />
  }
}
