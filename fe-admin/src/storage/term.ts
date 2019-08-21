/** 学期 */
import AWStorage from 'aw-easy-storage';
import createStorage from '@/storage/base';

class TermStorage {

  private static _instance: TermStorage | null = null;
  private storage: AWStorage;

  public static instance(): TermStorage {
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

  public key = 'termid';

  public clean() {
    return this.storage.remove(this.key);
  }

  public get(): number | null {
    return this.storage.getNumber(this.key);
  }

  public set(id: number) {
    this.storage.setNumber(this.key, id);
  }

}

export default TermStorage.instance();

(window as any).__TermStorage__ = TermStorage.instance();
