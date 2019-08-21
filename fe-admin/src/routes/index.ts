import { PATCH_CALLBACK_ROUTE_MIDDLEWARE } from '@/config';
import { routes } from "@/routes/routes";
import { RouteMeta } from '@/routes/type';
import AWRouter, { AWRouteState } from 'aw-react-router';
import { loginRequiredMiddleware, guestMiddleware, authCheckMiddleware } from './middleware/auth';


// 加载路由配置
const awRouter = AWRouter.instance();
awRouter.load(routes, {
  // mode
  mode: 'hash',
  // not found route name
  notFoundRouteName: '404',
  // 全局路由中间件
  middlewares: [
    // 设置 title
    (state: AWRouteState<RouteMeta>) => {
      const meta = state.meta || {};
      const routeTitle = meta.title;

      if (routeTitle) {
        document.title = routeTitle;
      } else {
        document.title = ['教学平台', meta.hanName].filter(Boolean).join('-');
      }

      return null;
    },
    PATCH_CALLBACK_ROUTE_MIDDLEWARE,                           // 暴露在 MAIN_CONFIG 中用于调试或打补丁的中间件
    loginRequiredMiddleware(['404', 'unauthorized', 'login']), // 这些路由不需验证是否登录
    guestMiddleware(['login']),                                // 只能未登录才可访问
    // 权限判断 (根据各路由配置的 meta.auth 进行判断)
    authCheckMiddleware(),
  ],
});

export default awRouter.renderRootRoute();
