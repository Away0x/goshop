import { useMedia } from 'use-media';

export function useMobile(): boolean {
  return useMedia({ maxWidth: 1000 });
}