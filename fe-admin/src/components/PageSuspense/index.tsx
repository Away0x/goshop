import React from 'react';
import { createSuspense } from 'aw-react-router';

export default createSuspense({
  defaultDelay: 0,
  defaultLoading: <div className="common_page_placeholder">loading ... ...</div>
});
