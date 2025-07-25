import { Outlet } from "react-router";

import { Header } from "@/components/header";

export function AppLayout() {
  return (
    <div className="antialised flex min-h-screen flex-col">
      <Header />

      <div className="flex flex-1 flex-col gap-4 p-8 pt-6">
        <Outlet />
      </div>
    </div>
  );
}
