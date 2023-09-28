import { useEffect, useState } from "react";
import useLocalStorage from "react-use/lib/useLocalStorage";
import { compare } from "semver";
import * as api from "@/helpers/api";
import { useGlobalStore } from "@/store/module";
import Icon from "./Icon";

interface State {
  latestVersion: string;
  show: boolean;
}

const UpgradeVersionView: React.FC = () => {
  const globalStore = useGlobalStore();
  const [skippedVersion, setSkippedVersion] = useLocalStorage<string>("skipped_version", "0.0.0");
  const profile = globalStore.state.systemStatus.profile;
  const [state, setState] = useState<State>({
    latestVersion: "",
    show: false,
  });

  useEffect(() => {
    api.getRepoLatestTag().then((latestTag) => {
      const latestVersion = latestTag.slice(1) || "0.0.0";
      const currentVersion = profile.version;
      const skipped = skippedVersion ? skippedVersion === latestVersion : false;
      setState({
        latestVersion,
        show: !skipped && compare(currentVersion, latestVersion) === -1,
      });
    });
  }, []);

  const onSkip = () => {
    setSkippedVersion(state.latestVersion);
    setState((s) => ({
      ...s,
      show: false,
    }));
  };

  if (!state.show) return null;

  return (
    <div className="flex flex-row justify-center items-center w-full py-2 px-2">
      <a
        className="flex flex-row justify-start items-center text-sm break-all text-green-600 hover:underline"
        target="_blank"
        href="https://github.com/usememos/memos/releases"
      >
        ✨ New version: v{state.latestVersion}
      </a>
      <button className="ml-1 opacity-60 text-gray-600 hover:opacity-100" onClick={onSkip}>
        <Icon.X className="w-4 h-auto" />
      </button>
    </div>
  );
};

export default UpgradeVersionView;
