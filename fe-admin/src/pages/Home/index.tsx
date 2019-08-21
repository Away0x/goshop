import React, { FC } from 'react';
import { Link } from 'aw-react-router';
import {
  Col,
} from 'antd';

import HomeHeader from './HomeHeader';
import HomeFooter from './HomeFooter';

import AWRouter, { AWRouteInfo } from 'aw-react-router';
import Authority from '@/auth/authority';
import { RouteMeta } from '@/routes/type';
import { isUndefined } from '@/tools/instance';

import './index.less';


const PAGE_NAMES = [
  {
    name: 'resource',
    img: require('@/assets/images/home/resource_module_new.png'),
  },
  {
    name: 'coach',
    img: require('@/assets/images/home/coach_module_new.png'),
  },
  {
    name: 'guide',
    img: require('@/assets/images/home/guide_module_new.png'),
  },
  {
    name: 'wrong',
    img: require('@/assets/images/home/wrong_module_new.png'),
  },
  {
    name: 'classroom',
    img: require('@/assets/images/home/class_module_new.png'),
  },
  {
    name: 'work',
    img: require('@/assets/images/home/work_module_new.png'),
  },
  {
    name: 'statistics',
    img: require('@/assets/images/home/statistics_module_new.png'),
  },
  {
    name: 'questions',
    img: require('@/assets/images/home/questions_module_new.png'),
  },
  {
    name: 'exam',
    img: require('@/assets/images/home/exam_module.png'),
  },
];


interface LinkNav {
  img: string,
  route: AWRouteInfo<RouteMeta>
}

const HomeContent: FC = () => {
  const navs: LinkNav[] = [];
  const routeInfos = AWRouter.instance().findMap(PAGE_NAMES.map(t => t.name));
  PAGE_NAMES.forEach((r) => {
    const name = r.name;
    const info: AWRouteInfo<RouteMeta> = routeInfos[name];
    if (!info) { return; }
    const meta = info.meta || {};
    const strict = isUndefined(meta.authUnstrict) ? true : !meta.authUnstrict;;

    if (Authority.checkAll(meta.auth || [], strict)) {
      navs.push({
        img: r.img,
        route: routeInfos[name],
      });
    }
  });

  return (
    <div className="home_page_content">
      <div className="before_line"></div>
      <div className="home_page__content clearfix">
        {
          navs.map((n) => {
            return (
              <Col span={5} offset={1} className="home_page__content_item" key={n.route.fullPath}>
                <Link to={n.route.fullPath}>
                  <img src={n.img} alt={n.route.name} />
                  <p>{n.route.meta.hanName}</p>
                </Link>
              </Col>
            );
          })
        }
      </div>
    </div>
  );
}

const HomePage: FC = () => {
  return (
    <div className="home_page">
      <HomeHeader />
      <HomeContent />
      <HomeFooter />
    </div>
  );
};

export default HomePage;
