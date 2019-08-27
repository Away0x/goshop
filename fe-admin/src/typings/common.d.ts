declare namespace AW {

  /** 路由 meta 信息 */
  export interface RouteMeta {
    readonly hanName: string // 路由的汉字标题
    readonly title: string // 页面的 title
    readonly topNavName?: string // 路由具体的 nav name，对应导航栏上的 nav name (高亮导航)
    readonly auth?: string[] // 权限数组
    // 权限非严格匹配(默认严格匹配)
    // - 严格匹配: 所给的 auth 数组中的所需权限都存在才有访问权限
    // - 非严格匹配: 所给的 auth 数组中的所需权限只要有其一即有访问权限
    readonly authUnstrict?: boolean;
  }

  /** js-cookie */
  export interface CookiesType {
    set(key: string, val: string, attr?: { [key: string]: any }): string | void;
    get(key: string, json?: boolean): any;
    remove(key: string, attr?: { [key: string]: any }): void;
  }

}
