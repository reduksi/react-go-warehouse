
import clsx from "clsx";
import { twMerge } from "tailwind-merge";

export const mc = (...classes) => twMerge(clsx(...classes));
