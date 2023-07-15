import { Divider } from "@mui/joy";
import { useEffect, useState } from "react";
import { toast } from "react-hot-toast";
import { useTranslate } from "@/utils/i18n";
import * as api from "@/helpers/api";
import showCreateIdentityProviderDialog from "../CreateIdentityProviderDialog";
import Dropdown from "../kit/Dropdown";
import { showCommonDialog } from "../Dialog/CommonDialog";
import LearnMore from "../LearnMore";

const SSOSection = () => {
  const t = useTranslate();
  const [identityProviderList, setIdentityProviderList] = useState<IdentityProvider[]>([]);

  useEffect(() => {
    fetchIdentityProviderList();
  }, []);

  const fetchIdentityProviderList = async () => {
    const { data: identityProviderList } = await api.getIdentityProviderList();
    setIdentityProviderList(identityProviderList);
  };

  const handleDeleteIdentityProvider = async (identityProvider: IdentityProvider) => {
    showCommonDialog({
      title: t("setting.sso-section.delete-sso"),
      content: t("setting.sso-section.confirm-delete", { name: identityProvider.name }),
      style: "warning",
      dialogName: "delete-identity-provider-dialog",
      onConfirm: async () => {
        try {
          await api.deleteIdentityProvider(identityProvider.id);
        } catch (error: any) {
          console.error(error);
          toast.error(error.response.data.message);
        }
        await fetchIdentityProviderList();
      },
    });
  };

  return (
    <div className="section-container">
      <div className="mb-2 w-full flex flex-row justify-start items-center gap-1">
        <span className="font-mono text-sm text-gray-400">{t("setting.sso-section.sso-list")}</span>
        <LearnMore url="https://usememos.com/docs/keycloak" />
        <button
          className="btn-normal px-2 py-0 ml-1"
          onClick={() => showCreateIdentityProviderDialog(undefined, fetchIdentityProviderList)}
        >
          {t("common.create")}
        </button>
      </div>

      <Divider />

      {identityProviderList.map((identityProvider) => (
        <div
          key={identityProvider.id}
          className="py-2 w-full border-b last:border-b dark:border-zinc-700 flex flex-row items-center justify-between"
        >
          <div className="flex flex-row items-center">
            <p className="ml-2">
              {identityProvider.name}
              <span className="text-sm ml-1 opacity-40">({identityProvider.type})</span>
            </p>
          </div>
          <div className="flex flex-row items-center">
            <Dropdown
              actionsClassName="!w-28"
              actions={
                <>
                  <button
                    className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded hover:bg-gray-100 dark:hover:bg-zinc-600"
                    onClick={() => showCreateIdentityProviderDialog(identityProvider, fetchIdentityProviderList)}
                  >
                    {t("common.edit")}
                  </button>
                  <button
                    className="w-full text-left text-sm leading-6 py-1 px-3 cursor-pointer rounded text-red-600 hover:bg-gray-100 dark:hover:bg-zinc-600"
                    onClick={() => handleDeleteIdentityProvider(identityProvider)}
                  >
                    {t("common.delete")}
                  </button>
                </>
              }
            />
          </div>
        </div>
      ))}
    </div>
  );
};

export default SSOSection;
