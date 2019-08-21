declare namespace AW {

  /** 用户信息 */
  export interface UserEntity {
    userid: number;   // 用户 id
    usercode: string; // user code
    username: string; // 用户名
    schoolid: number; // 学校 id
    source?: number;  // 来源 id
    category: 0 | 1 | 2 | 3; // 分类 id

    address?: string;   // 地址
    birthday?: string;  // 生日
    created_at: number; // 创建时间
    email: string;      // 邮箱
    gender: number;     // 性别
    tel: string;        // 电话
    photo?: string;     // 头像
    fullPath?: string;  // 头像图片地址
    classid?: number;   // 班级id
    loginname?: string; // 用户别名

    acl?: UserAuthMapType; // 权限映射表
    aclrole?: UserRoleType[]; // 角色信息

    // true: 包含角色不需要跳转到设置班级页面
    // false: 没有对应角色信息，需跳转到班级设置页面
    hasRole?: boolean;
    // 是否有设置班级
    hasClassinfo?: boolean;
  }

  /** 角色信息 */
  export interface UserRoleType {
    activestatus: 0 | 1, // 1 为默认角色
    name: string,
    roleid: number, // 角色标识
    type: number,
  }

  /** 权限映射表 */
  export interface UserAuthMapType {
    [key: string]: string[];
  }

  /** 携带学期信息的用户信息 */
  export interface UserEntityAndTermInfoType {
    userEntity: UserEntity;
    termVos: TermType[];
  }

  /** 用户信息和学期数据 */
  export interface AuthAPIResponseType {
    userEntity: UserEntity;
    termVos: TermType[];
  }

}