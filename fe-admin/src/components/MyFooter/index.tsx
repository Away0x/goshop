import React, { FC } from 'react';
import {
  Row,
  Col,
} from 'antd';

import './index.less';

const year = new Date().getFullYear();

const MyFooter: FC = () => {
  return (
    <footer className="my_footer_component">
      <Row className="inner">
        <Col span={10}>Copyright © 2009 - {year} Edutech, Inc. All rights reserved.</Col>
        <Col span={14}>
          版权所有：上海易教信息科技有限公司
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          <a rel="noopener noreferrer" href="http://www.educationtek.com/" target="_blank">www.educationtek.com</a>
        </Col>
      </Row>
    </footer>
  );
}

export default MyFooter;
