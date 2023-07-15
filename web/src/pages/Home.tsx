import { useEffect } from "react";
import { toast } from "react-hot-toast";
import { useTranslate } from "@/utils/i18n";
import { useGlobalStore, useUserStore } from "@/store/module";
import MemoEditor from "@/components/MemoEditor";
import MemoFilter from "@/components/MemoFilter";
import MemoList from "@/components/MemoList";
import MobileHeader from "@/components/MobileHeader";
import HomeSidebar from "@/components/HomeSidebar";

function Home() {
  const t = useTranslate();
  const globalStore = useGlobalStore();
  const userStore = useUserStore();
  const user = userStore.state.user;

  useEffect(() => {
    const currentUserId = userStore.getCurrentUserId();
    userStore.getUserById(currentUserId).then((user) => {
      if (!user) {
        toast.error(t("message.user-not-found"));
        return;
      }
    });
  }, [userStore.getCurrentUserId()]);

  useEffect(() => {
    if (user?.setting.locale) {
      globalStore.setLocale(user.setting.locale);
    }
  }, [user?.setting.locale]);

  return (
    <div className="w-full flex flex-row justify-start items-start">
      <div className="flex-grow shrink w-auto px-4 sm:px-2 sm:pt-4">
        <MobileHeader />
        <div className="w-full h-auto flex flex-col justify-start items-start bg-zinc-100 dark:bg-zinc-800 rounded-lg">
          {!userStore.isVisitorMode() && <MemoEditor className="mb-2" />}
          <MemoFilter />
        </div>
        <MemoList />
      </div>
      <HomeSidebar />
    </div>
  );
}

export default Home;
