import React, { FC } from 'react';
import {
  Row,
  Col,
} from 'antd';

import ChapterKnowTree from '@/components/ChapterKnowTree';

const LeftTreeLayout: FC = (props) => {
  return (
    <Row className="left_tree_layout">
      <Col span={4}>
        <ChapterKnowTree />
      </Col>
      <Col span={20}>
        {props.children}
      </Col>
    </Row>
  );
};

export default LeftTreeLayout;
