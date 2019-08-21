import React, { FC, memo } from 'react';

import { useBreadcrumbs } from '@/layouts/Base/hooks';
import LeftTreeLayout from '@/layouts/LeftTree';

const WorkListPage: FC = () => {
  useBreadcrumbs([{name: '网络作业'}]);

  return (
    <LeftTreeLayout>
      <div>work list</div>
    </LeftTreeLayout>
  );
};

export default memo(WorkListPage);
