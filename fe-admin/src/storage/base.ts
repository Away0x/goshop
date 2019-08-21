import AWStorage, { AWStorageConstructorOptions, NO_VALIDATE_EXPIRE } from 'aw-easy-storage';
import { LOCALSTROAGE_PREFIX, DISABLED_EXPIRE } from '@/config';

const AWStorageFactory = ({
  enableLog = false,
  disabled = false,
  prefix = LOCALSTROAGE_PREFIX,
  session = false,
  hasMeta = true,
  expire = NO_VALIDATE_EXPIRE,
  // 当 expire 设置为 DISABLED_EXPIRE，则相当于永远都是过期的 (禁用缓存)
  checkExpire = (e) => e === DISABLED_EXPIRE,
}: Partial<AWStorageConstructorOptions>) => {
  return new AWStorage({
    enableLog,
    disabled,
    session,
    prefix,
    hasMeta,
    expire,
    checkExpire,
  });
};

export default AWStorageFactory;
