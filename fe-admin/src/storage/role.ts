/** 角色 */
import AWStorage from 'aw-easy-storage';
import createStorage from '@/storage/base';

interface RoleType {
  roleid: number
  roletype: number
}

class RoleStorage {

  private static _instance: RoleStorage | null = null;
  private storage: AWStorage;

  public static instance(): RoleStorage {
    if (this._instance === null) {
      this._instance = new this();
    }

    return this._instance!;
  }

  constructor() {
    this.storage = createStorage({
      hasMeta: false,
    });
  }

  private roleStore: RoleType | null = null;
  public key = 'role';

  public clean() {
    this.roleStore = null;
    return this.storage.remove(this.key);
  }

  public get(): RoleType | null {
    if (this.roleStore) {
      return this.roleStore;
    }

    try {
      const jsonString = this.storage.getString(this.key);
      if (!jsonString) {
        return null;
      }
      const data = JSON.parse(jsonString);
      if (!data.roleid || !data.roletype) {
        return null;
      }
      const roleData: RoleType = {
        roleid: Number(data.roleid),
        roletype: Number(data.roletype),
      };
      this.roleStore = roleData;

      return roleData;
    } catch(err) {
      console.warn(err);
      return null;
    }
  }

  public set(role: RoleType) {
    this.roleStore = role;
    this.storage.setString(this.key, JSON.stringify(role));
  }

  public isTeacherRole(): boolean {
    const role = this.get();
    if (!role) {
      return true; // 默认老师
    }
    // 1 为教师角色 (注意: 这里的 type 和用户 category 不一样，那个是判断学生老师管理员身份的)
    return role.roletype === 1;
  }

}

export default RoleStorage.instance();

(window as any).__RoleStorage__ = RoleStorage.instance();
