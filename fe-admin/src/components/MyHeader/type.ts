import { AWRouteInfo } from "aw-react-router";
import { RouteMeta } from "@/routes/type";

export interface Nav {
  defaultName?: string
  subNav?: Nav[]
  route?: AWRouteInfo<RouteMeta>
}
