import { useEffect } from "react";
import { toast } from "react-hot-toast";
import { useTranslate } from "@/utils/i18n";
import { useFilterStore, useShortcutStore } from "@/store/module";
import { getTimeStampByDate } from "@/helpers/datetime";
import useToggle from "@/hooks/useToggle";
import useLoading from "@/hooks/useLoading";
import Icon from "./Icon";
import showCreateShortcutDialog from "./CreateShortcutDialog";

const ShortcutList = () => {
  const t = useTranslate();
  const filterStore = useFilterStore();
  const shortcutStore = useShortcutStore();
  const filter = filterStore.state;
  const shortcuts = shortcutStore.state.shortcuts;
  const loadingState = useLoading();

  const pinnedShortcuts = shortcuts
    .filter((s) => s.rowStatus === "ARCHIVED")
    .sort((a, b) => getTimeStampByDate(b.createdTs) - getTimeStampByDate(a.createdTs));
  const unpinnedShortcuts = shortcuts
    .filter((s) => s.rowStatus === "NORMAL")
    .sort((a, b) => getTimeStampByDate(b.createdTs) - getTimeStampByDate(a.createdTs));
  const sortedShortcuts = pinnedShortcuts.concat(unpinnedShortcuts);

  useEffect(() => {
    shortcutStore
      .getMyAllShortcuts()
      .catch(() => {
        // do nth
      })
      .finally(() => {
        loadingState.setFinish();
      });
  }, []);

  return (
    <div className="flex flex-col justify-start items-start w-full mt-2 h-auto shrink-0 flex-nowrap hide-scrollbar">
      <div className="flex flex-row justify-start items-center w-full px-4">
        <span className="text-sm leading-6 font-mono text-gray-400">{t("common.shortcuts")}</span>
        <button
          className="flex flex-col justify-center items-center w-5 h-5 bg-gray-200 dark:bg-zinc-700 rounded ml-2 hover:shadow"
          onClick={() => showCreateShortcutDialog()}
        >
          <Icon.Plus className="w-4 h-4 text-gray-400" />
        </button>
      </div>
      <div className="flex flex-col justify-start items-start relative w-full h-auto flex-nowrap mb-2">
        {sortedShortcuts.map((s) => {
          return <ShortcutContainer key={s.id} shortcut={s} isActive={s.id === Number(filter?.shortcutId)} />;
        })}
      </div>
    </div>
  );
};

interface ShortcutContainerProps {
  shortcut: Shortcut;
  isActive: boolean;
}

const ShortcutContainer: React.FC<ShortcutContainerProps> = (props: ShortcutContainerProps) => {
  const { shortcut, isActive } = props;
  const t = useTranslate();
  const filterStore = useFilterStore();
  const shortcutStore = useShortcutStore();
  const [showConfirmDeleteBtn, toggleConfirmDeleteBtn] = useToggle(false);

  const handleShortcutClick = () => {
    if (isActive) {
      filterStore.setMemoShortcut(undefined);
    } else {
      filterStore.setMemoShortcut(shortcut.id);
    }
  };

  const handleDeleteMemoClick = async (event: React.MouseEvent) => {
    event.stopPropagation();

    if (showConfirmDeleteBtn) {
      try {
        await shortcutStore.deleteShortcutById(shortcut.id);
        if (filterStore.getState().shortcutId === shortcut.id) {
          // need clear shortcut filter
          filterStore.setMemoShortcut(undefined);
        }
      } catch (error: any) {
        console.error(error);
        toast.error(error.response.data.message);
      }
    } else {
      toggleConfirmDeleteBtn();
    }
  };

  const handleEditShortcutBtnClick = (event: React.MouseEvent) => {
    event.stopPropagation();
    showCreateShortcutDialog(shortcut.id);
  };

  const handlePinShortcutBtnClick = async (event: React.MouseEvent) => {
    event.stopPropagation();

    try {
      const shortcutPatch: ShortcutPatch = {
        id: shortcut.id,
        rowStatus: shortcut.rowStatus === "ARCHIVED" ? "NORMAL" : "ARCHIVED",
      };
      await shortcutStore.patchShortcut(shortcutPatch);
    } catch (error) {
      // do nth
    }
  };

  const handleDeleteBtnMouseLeave = () => {
    toggleConfirmDeleteBtn(false);
  };

  return (
    <>
      <div
        className="relative group flex flex-row justify-between items-center w-full h-10 py-0 px-4 mt-px first:mt-2 rounded-lg text-base cursor-pointer select-none shrink-0 hover:bg-white dark:hover:bg-zinc-700"
        onClick={handleShortcutClick}
      >
        <div
          className={`flex flex-row justify-start items-center truncate shrink leading-5 mr-1 text-black dark:text-gray-200 ${
            isActive && "text-green-600"
          }`}
        >
          <span className="truncate">{shortcut.title}</span>
        </div>
        <div className="flex-row justify-end items-center hidden group/btns group-hover:flex shrink-0">
          <span className="flex flex-row justify-center items-center toggle-btn">
            <Icon.MoreHorizontal className="w-4 h-auto" />
          </span>
          <div className="absolute top-4 right-0 flex-col justify-start items-start w-auto h-auto px-4 pt-3 hidden group-hover/btns:flex z-1">
            <div className="flex flex-col justify-start items-start w-32 h-auto p-1 whitespace-nowrap rounded-md bg-white dark:bg-zinc-700 shadow">
              <span
                className="w-full text-sm leading-6 py-1 px-3 rounded text-left dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-zinc-800"
                onClick={handlePinShortcutBtnClick}
              >
                {shortcut.rowStatus === "ARCHIVED" ? t("common.unpin") : t("common.pin")}
              </span>
              <span
                className="w-full text-sm leading-6 py-1 px-3 rounded text-left dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-zinc-800"
                onClick={handleEditShortcutBtnClick}
              >
                {t("common.edit")}
              </span>
              <span
                className={`w-full text-sm leading-6 py-1 px-3 rounded text-left dark:text-gray-200 hover:bg-gray-100 dark:hover:bg-zinc-800 text-orange-600 ${
                  showConfirmDeleteBtn && "font-black"
                }`}
                onClick={handleDeleteMemoClick}
                onMouseLeave={handleDeleteBtnMouseLeave}
              >
                {t("common.delete")}
                {showConfirmDeleteBtn ? "!" : ""}
              </span>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default ShortcutList;
