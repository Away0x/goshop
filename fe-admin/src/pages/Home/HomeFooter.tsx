import React, { FC } from 'react';

const HomeFooter: FC = () => {
  return (
    <footer>
      <div className="after_line"></div>
      <div className="inner">
        <p>Copyright © 2009 - {new Date().getFullYear()} Edutech, Inc. All rights reserved.</p>
        <p>版权所有：上海易教信息科技有限公司
          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <a rel="noopener noreferrer" href="http://www.educationtek.com/" target="_blank">www.educationtek.com</a>
        </p>
      </div>
    </footer>
  );
}

export default HomeFooter;
