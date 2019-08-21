import { useEffect } from 'react';

type Callback = (...args: any[]) => void | any;

export function useOnUnmount(onUnmount: Callback) {
  return useEffect(() => {
    return () => onUnmount && onUnmount();
  }, []);
}

export function useOnMount(onMount: Callback) {
  return useEffect(() => {
    onMount && onMount();
  }, []);
}
  
export function useOnUpdate(onUpdate: Callback) {
  return useEffect(() => {
    onUpdate && onUpdate();
  });
}