// ---------------------------------------- css ----------------------------------------
import '@/assets/styles/index.less';

// ---------------------------------------- main ----------------------------------------
import React from 'react';
import ReactDOM from 'react-dom';
import FastClick from 'fastclick';
import { message as messageAlert } from 'antd';

// ---------------------------------------- 其他 ----------------------------------------
import { APP_ROOT_ID } from '@/constants';
import { isDev } from "@/config";
import App from '@/App';
import Token from '@/auth/token';

// ---------------------------------------- 错误处理 ----------------------------------------
window.onerror = (event: Event | string, source?: string, fileno?: number, columnNumber?: number, error?: Error) => {
  const error_data = {
    title: document.getElementsByTagName('title')[0].innerText,
    error_event: event,
    script_url: source,
    line_number: fileno,
    column_number: columnNumber,
    error_detail: (error && error.message) || '',
    error_stack: (error && error.stack) || '',
    user_agent: window.navigator.userAgent,
    location_href: window.location.href,
    token: Token.get(),
  };
  const event_string = String(error_data.error_event || '');

  if (event_string === 'ResizeObserver loop limit exceeded') { return; }

  const SCRIPT_ERROR = 'Unexpected token <'; // 脚本错误 (一般是服务器中的前端代码更新了，浏览器缓存中的 js 代码由于不一致而报错)
  const CHUNK_ERROR = 'chunk';

  if (event_string.indexOf(SCRIPT_ERROR) !== -1 || event_string.indexOf(CHUNK_ERROR) !== -1) {
    messageAlert.warning('发现新版本，请使用刷新网页');
  }
};

// ---------------------------------------- 挂载 ----------------------------------------
(FastClick as any).attach(document.body);

// ---------------------------------------- token ----------------------------------------
window.addEventListener('storage', (ev: StorageEvent) => {
  const key = ev.key;
  const val = ev.newValue;

  if (key === Token.key && val === null) { Token.clean(); }
});

// logout 时的 callback
Token.logoutCallback = () => {
  window.localStorage.clear();
  window.sessionStorage.clear();
}

// ======================================= RENDER =======================================
function createApp() {
  return (
    <App />
  );
}

function render() {
  ReactDOM.render(
    createApp(),
    document.getElementById(APP_ROOT_ID),
  );
}

if (isDev()) {
  // DEVELOPMENT
  const devRender = () => {
    if ((module as any).hot) {
      (module as any).hot.accept('./App', () => render());
    }

    render();
  }

  try {
    devRender();
  } catch (error) {
    console.error(error);
    render();
  }
} else {
  // PRODUCTION
  render();
}
