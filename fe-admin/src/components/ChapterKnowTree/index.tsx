import React, { FC, memo } from 'react';
// import {
//   Row,
//   Col,
// } from 'antd';

import { withDefaultProps } from '@/tools/hoc';

// import MyTree from '@/components/MyTree';
import TreeTypeSwitch from './TreeTypeSwitch';

// enum TreeType {
//   CHAPTER = 'chapter', // 章节树
//   KNOW = 'know'        // 知识点树
// }

interface ChapterKnowTreeProps {

}

const defaultProps = {
  height: 620,
  hasKnowTree: true,
  hasFilter: true,
};

const ChapterKnowTree: FC<(typeof defaultProps) & ChapterKnowTreeProps> = () => {
  return (
    <div className="chapter_know_tree_component">

      {/* tree type switch */}
      <TreeTypeSwitch />
    </div>
  );
};

export default withDefaultProps(defaultProps, memo(ChapterKnowTree));
