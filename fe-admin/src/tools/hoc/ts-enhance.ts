// ts react 的一些推断辅助
import { ComponentType } from "react";

type Omit<T, K extends keyof any> = Pick<T, Exclude<keyof T, K>>;

/**
 * typescript defaultProps 的默认 props 辅助 hoc
 */
export const withDefaultProps = <P extends object, DP extends Partial<P> = Partial<P>>(
  defaultProps: DP,
  Cmp: ComponentType<P>,
) => {
  type RequiredProp = Omit<P, keyof DP>;
  type Props = Partial<DP> & Required<RequiredProp>;

  Cmp.defaultProps = defaultProps;

  return (Cmp as ComponentType<any>) as ComponentType<Props>;
};
