import copy from "copy-to-clipboard";
import React from "react";
import toast from "react-hot-toast";
import { useTranslate } from "@/utils/i18n";
import { useResourceStore } from "@/store/module";
import { getResourceUrl } from "@/utils/resource";
import Dropdown from "./kit/Dropdown";
import Icon from "./Icon";
import { showCommonDialog } from "./Dialog/CommonDialog";
import showChangeResourceFilenameDialog from "./ChangeResourceFilenameDialog";
import showPreviewImageDialog from "./PreviewImageDialog";

interface Props {
  resource: Resource;
}

const ResourceItemDropdown = ({ resource }: Props) => {
  const t = useTranslate();
  const resourceStore = useResourceStore();

  const handlePreviewBtnClick = (resource: Resource) => {
    const resourceUrl = getResourceUrl(resource);
    if (resource.type.startsWith("image")) {
      showPreviewImageDialog([getResourceUrl(resource)], 0);
    } else {
      window.open(resourceUrl);
    }
  };

  const handleCopyResourceLinkBtnClick = (resource: Resource) => {
    const url = getResourceUrl(resource);
    copy(url);
    toast.success(t("message.succeed-copy-resource-link"));
  };

  const handleRenameBtnClick = (resource: Resource) => {
    showChangeResourceFilenameDialog(resource.id, resource.filename);
  };

  const handleDeleteResourceBtnClick = (resource: Resource) => {
    let warningText = t("resource.warning-text");
    if (resource.linkedMemoAmount > 0) {
      warningText = warningText + `\n${t("resource.linked-amount")}: ${resource.linkedMemoAmount}`;
    }

    showCommonDialog({
      title: t("resource.delete-resource"),
      content: warningText,
      style: "warning",
      dialogName: "delete-resource-dialog",
      onConfirm: async () => {
        await resourceStore.deleteResourceById(resource.id);
      },
    });
  };

  return (
    <Dropdown
      actionsClassName="!w-auto min-w-[8rem]"
      trigger={<Icon.MoreVertical className="w-4 h-auto hover:opacity-80 cursor-pointer" />}
      actions={
        <>
          <button
            className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded hover:bg-gray-100 dark:hover:bg-zinc-600"
            onClick={() => handlePreviewBtnClick(resource)}
          >
            {t("common.preview")}
          </button>
          <button
            className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded hover:bg-gray-100 dark:hover:bg-zinc-600"
            onClick={() => handleCopyResourceLinkBtnClick(resource)}
          >
            {t("resource.copy-link")}
          </button>
          <button
            className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded hover:bg-gray-100 dark:hover:bg-zinc-600"
            onClick={() => handleRenameBtnClick(resource)}
          >
            {t("common.rename")}
          </button>
          <button
            className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded text-red-600 hover:bg-gray-100 dark:hover:bg-zinc-600"
            onClick={() => handleDeleteResourceBtnClick(resource)}
          >
            {t("common.delete")}
          </button>
        </>
      }
    />
  );
};

export default React.memo(ResourceItemDropdown);
