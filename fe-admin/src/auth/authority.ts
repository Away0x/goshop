// 权限模块
interface AuthorityConstructorOptions {
  /** 是否 mock 权限 (会得到所有权限) */
  mock: boolean
}

class Authority {

  private static _instance: Authority | null = null;

  public static instance(opt?: AuthorityConstructorOptions): Authority {
    if (this._instance === null) {
      this._instance = new this(opt || {});
    }

    return this._instance!;
  }

  /** 不判断权限 (用于调试) */
  private mock: boolean;

  constructor({
   mock = false,
  }: Partial<AuthorityConstructorOptions>) {
    this.mock = mock;
  }

  /** 是否已加载过权限 */
  public loaded = false;
  /** 权限表 */
  public authStore: string[] = [];

  /** 加载权限 */
  public load(authArr: string[] = []) {
    if (this.loaded) {
      return;
    }

    this.authStore = authArr;
    this.loaded = true;
  }

  /** 加载权限 (map 形式) */
  public loadMap(authMap: AW.UserAuthMapType = {}) {
    if (this.loaded) {
      return;
    }

    let authArr: string[] = [];
    for (const key in authMap) {
      if (authMap.hasOwnProperty(key)) {
        const item = authMap[key];
        authArr.push(key);
        authArr = authArr.concat(item);
      }
    }

    this.load(authArr);
  }

  /** 重新加载权限 */
  public reload(authArr: string[] = []) {
    this.authStore = authArr;
    this.loaded = true;
  }

  /** 检查权限 */
  public check(auth: string): boolean {
    if (!this.loaded) {
      throw new Error('Authority: 还未加载权限数据');
    }

    if (this.mock) { return true; }

    if (!this.authStore.length) {
      return false;
    }

    return this.authStore.indexOf(auth) !== -1;
  }

  /**
   * 检查权限
   * strict: true，auths 中必须全部存在才通过，否则一个通过即可
   */
  public checkAll(auths: string[], strict = true): boolean {
    if (strict) {
      return auths.every((a) => this.check(a));
    }

    return auths.some((a) => this.check(a));
  }

}

const instance = Authority.instance({
  mock: false,
});

export default instance;

(window as any).__Authority__ = instance;
