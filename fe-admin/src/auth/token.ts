// 登出、获取用户 token 操作
import TokenStorage from '@/storage/token';

interface LogoutConfig {
  beforeCallback?: (token: string) => Promise<boolean> // 当前退出行为前的 callback
  afterCallback?: () => Promise<boolean>               //当前退出行为后的 callback
  triggerGlobalCallback?: boolean                      // 触发全局的 logout callback
  reloadPage?: boolean                                 // 是否要刷新页面
}

export default class {

  public static readonly key = TokenStorage.key;
  // main.ts 中挂载的回调 (将不方便在这个类中做的事情放到 main.ts 中去做)
  public static logoutCallback: null | (() => void) = null;

  public static set(token: string) {
    TokenStorage.set(token);
  }

  public static get(): string {
    return TokenStorage.get();
  }

  public static isLogin(): boolean {
    return !!this.get();
  }

  public static clean() {
    TokenStorage.clean();
  }

  /**
   * 登出操作
   */
  public static async logout(config: LogoutConfig) {
    if (config.beforeCallback) {
      const status = await config.beforeCallback(this.get());
      if (!status) {
        return;
      }
    }

    this.clean();

    if (config.afterCallback) {
      const status = await config.afterCallback();
      if (!status) {
        return;
      }
    }

    // 触发全局的 callback
    if (config.triggerGlobalCallback && this.logoutCallback) {
      await this.logoutCallback();
    }

    if (config.reloadPage) {
      window.location.href = '/';
    }
  }

}
