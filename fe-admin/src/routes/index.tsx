import React from 'react';
import AWRouter, { AWRouteConfig } from 'aw-react-router';

import PageSuspense from '@/components/PageSuspense';
import { PATCH_CALLBACK_ROUTE_MIDDLEWARE } from '@/config';
// import { loginRequiredMiddleware, guestMiddleware, authCheckMiddleware } from '@/routes/middleware/auth';

const routes: AWRouteConfig<AW.RouteMeta>[] = [
  {
    name: 'login',
    path: '/login',
    component: () => <div>login</div>,
    meta: {
      title: '登录',
      hanName: '登录',
    },
  },

  {
    name: '404',
    path: '/404',
    component: (props) => <PageSuspense {...props} load={import(/* webpackChunkName: 'notfound' */ '@/pages/Errors/NotFound')} />,
    meta: {
      title: '404',
      hanName: '页面未找到',
    },
  },
  {
    name: 'unauthorized',
    path: '/unauthorized',
    component: (props) => <PageSuspense {...props} load={import(/* webpackChunkName: 'unauthorized' */ '@/pages/Errors/Unauthorized')} />,
    meta: {
      title: 'unauthorized',
      hanName: '无权限访问',
    },
  },
];


// 加载路由配置
export const awRouter = AWRouter.instance<AW.RouteMeta>();
awRouter.load(routes, {
  mode: 'history',          // mode
  notFoundRouteName: '404', // not found route name
  // 全局路由中间件
  middlewares: [
    (state) => {
      const meta = state.meta || {};
      document.title = meta.title
        ? meta.title
        : ['管理员后台', meta.hanName].filter(Boolean).join('-');
      return null;
    },
    PATCH_CALLBACK_ROUTE_MIDDLEWARE,                           // 暴露在 MAIN_CONFIG 中用于调试或打补丁的中间件
    // loginRequiredMiddleware(['404', 'unauthorized', 'login']), // 这些路由不需验证是否登录
    // guestMiddleware(['login']),                                // 只能未登录才可访问
    // authCheckMiddleware(),                                     // 权限判断 (根据各路由配置的 meta.auth 进行判断)
  ],
});

export default awRouter.renderRootRoute();
