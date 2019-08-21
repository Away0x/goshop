// 用户相关常量

// 用户类型
export enum USER_TYPE {
  SUPER_ADMIN = 0, // 超级管理员
  ADMIN = 1,       // 学校管理员
  TEACHER = 2,     // 教师
  STUDENT = 3,     // 学生
}

/**
 * 角色类型
 * - 1: 教师 TEACHER
 * - 2: 班主任 HEAD_TEACHER
 * - 3: 学生 STUDENTS
 * - 4: 校长 PRINCIPAL
 * - 5: 年级主任 GRADE_DIRECTOR
 * - 6: 科目组长 SUBJECT_LEADER
 */
export enum ROLE_TYPE {
  TEACHER = 1,
  HEAD_TEACHER = 2,
  STUDENTS = 3,
  PRINCIPAL = 4,
  GRADE_DIRECTOR = 5,
  SUBJECT_LEADER = 6,
}

// 是否为默认角色
export enum ROLE_STATUS {
  IS = 1,
  NO = 0,
}
