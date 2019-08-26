import { useEffect } from 'react';

import BaseLayoutStore, { BreadcrumbsType } from '@/store/base-layout';
import Authority from '@/auth/authority';
import { isUndefined } from '@/tools/instance';
import { PAGE_NAMES } from '@/constants';
import { Nav } from '@/components/MyHeader/type';
import { awRouter } from '@/routes';

/** 修改面包屑 */
export function useBreadcrumbs(initBreadcrumbs: BreadcrumbsType[]) {
  const store = BaseLayoutStore.useStore();

  const setBreadcrumbs = (breadcrumbs: BreadcrumbsType[]) => {
    store.setBreadcrumbsAction([
      {
        name: '首页',
        routeName: 'home',
      },
      ...breadcrumbs,
    ]);
  };

  useEffect(() => {
    setBreadcrumbs(initBreadcrumbs);
  }, []);

  return {
    breadcrumbs: store.state.breadcrumbs,
    setBreadcrumbs,
  };
}


const NAV_MAX_LEN = 8; // 最后一个为 "更多"
/** 创建页面导航 */
export function useNavs(): Nav[] {
  const navs: Nav[] = [];
  const moreNavs: Nav = {
    defaultName: '更多',
    subNav: [],
  };
  const routeInfos = awRouter.findMap(PAGE_NAMES.map(t => t.name));

  PAGE_NAMES.forEach((r) => {
    const name = r.name;
    const info = routeInfos[name];
    if (!info) { return; }
    const meta = info.meta || {};
    const strict = isUndefined(meta.authUnstrict) ? true : !meta.authUnstrict;;

    if (Authority.checkAll(meta.auth || [], strict)) {
      if (navs.length >= (NAV_MAX_LEN - 1)) {
        moreNavs.subNav!.push({ route: routeInfos[name] });
        return;
      }
      navs.push({ route: routeInfos[name] });
    }
  });

  if (moreNavs.subNav!.length > 1) {
    return [...navs, moreNavs];
  } else if (moreNavs.subNav!.length === 1) {
    return [...navs, moreNavs.subNav![0]];
  }

  return navs;
}
