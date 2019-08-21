declare namespace AW {

  /** 学期 */
  export interface TermType {
    termid: number;     // 学期 id
    termname: string;   // 学期名称
    iscurrent?: number; // 是否为当前学期: 1 : 是当前学期; 0 : 不是当前学期
  }
  
}