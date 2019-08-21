/** token */
import AWStorage from 'aw-easy-storage';
import createStorage from '@/storage/base';
import { TOKEN_KEY, getCookie } from '@/config';

// 浏览器退出时清除 token
class TokenStorage {

  private static _instance: TokenStorage | null = null;
  private storage: AWStorage;

  public static instance(): TokenStorage {
    if (this._instance === null) {
      this._instance = new this();
    }

    return this._instance!;
  }

  constructor() {
    this.storage = createStorage({
      hasMeta: false,
      session: true,
      prefix: '',
    });
  }

  public readonly key = TOKEN_KEY;
  private tokenStore: null | string = null; // 内存缓存

  public set(token: string) {
    const cookie = getCookie();

    if (cookie) {
      cookie.set(this.key, token);
    }
    this.storage.setString(this.key, token);

    this.tokenStore = token;
  }

  /** 取值优先级 (内存 > cookie > sessionStorage) */
  public get(): string {
    if (this.tokenStore) {
      return this.tokenStore;
    }
    const cookie = getCookie();

    // cookie 有则读 cookie，否则读 sessionStorage，然后同步到 cookie
    let tokenFromLocal = '';
    if (cookie) { tokenFromLocal = cookie.get(this.key); }
    if (!tokenFromLocal) {
      tokenFromLocal = this.storage.getString(this.key) || '';
      if (cookie) { cookie.set(this.key, tokenFromLocal); }
    }

    this.tokenStore = tokenFromLocal;
    return this.tokenStore || '';
  }

  public clean() {
    const cookie = getCookie();

    if (cookie) { cookie.remove(this.key); }

    this.storage.remove(this.key);
    this.tokenStore = null;
  }

}

export default TokenStorage.instance();

(window as any).__TokenStorage__ = TokenStorage.instance();
