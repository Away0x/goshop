import React, { FC } from 'react';
import {
  Row,
  Col,
} from 'antd';

const TreeTypeSwitch: FC = () => {
  return (
    <Row className="chapter_know_tree__tree_type_switch">
      <Col span={12}>章节</Col>
      <Col span={12}>知识点</Col>
    </Row>
  );
};

export default TreeTypeSwitch;
