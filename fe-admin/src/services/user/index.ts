import Services from '@/services/base';
import { CommonResponse } from '@/services/type';
import { AUTH_API_ROOT } from '@/config';

interface LoginParams {
  username: string
  password: string
}

export default class {

  private static authAPI(api: string) {
    return AUTH_API_ROOT + api;
  };

  /** 登录 */
  public static async login({username, password}: LoginParams): Promise<CommonResponse<AW.AuthAPIResponseType | null>> {
    const resp = await Services.get<AW.AuthAPIResponseType>({
      url: this.authAPI('/efficiencyLogin'),
      params: { username, password },
      showErrorMessage: false,
    })

    return resp;
  }

  /** 获取用户信息 */
  public static async getUserInfo(): Promise<CommonResponse<AW.UserEntityAndTermInfoType | null>> {
    const resp = await Services.get<AW.UserEntityAndTermInfoType>({
      url: this.authAPI('/api/user/efficiency-select-one'),
    });

    return resp;
  }

  /** 获取默认学期 */
  public static async getDetaultTerm(): Promise<CommonResponse<AW.TermType>> {
    const resp = await Services.get<any>({
      url: this.authAPI('/api/term/get-default-terminfo'),
    });

    const _data = resp.data || {};
    const data: AW.TermType = {
      termid: _data.termid,
      termname: _data.yearname + '学年' + _data.termname,
    };

    return {
      status: resp.status,
      message: resp.message,
      data,
    };
  }

}
