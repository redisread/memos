# memos

<img height="72px" src="https://usememos.com/logo.png" alt="✍️ memos" align="right" />

A privacy-first, lightweight note-taking service. Easily capture and share your great thoughts.

<a href="https://www.usememos.com">Home Page</a> •
<a href="https://www.usememos.com/blog">Blogs</a> •
<a href="https://www.usememos.com/docs">Docs</a> •
<a href="https://demo.usememos.com/">Live Demo</a>

<p>
  <a href="https://github.com/usememos/memos/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/usememos/memos?logo=github" /></a>
  <a href="https://hub.docker.com/r/neosmemo/memos"><img alt="Docker pull" src="https://img.shields.io/docker/pulls/neosmemo/memos.svg"/></a>
  <a href="https://hosted.weblate.org/engage/memos-i18n/"><img src="https://hosted.weblate.org/widget/memos-i18n/english/svg-badge.svg" alt="Translation status" /></a>
  <a href="https://discord.gg/tfPJa4UmAv"><img alt="Discord" src="https://img.shields.io/badge/discord-chat-5865f2?logo=discord&logoColor=f5f5f5" /></a>
</p>

![demo](https://www.usememos.com/demo.webp)

## Key points

- **Open source and free forever**. Embrace a future where creativity knows no boundaries with our open-source solution – free today, tomorrow, and always.
- **Self-hosting with Docker in just seconds**. Enjoy the flexibility, scalability, and ease of setup that Docker provides, allowing you to have full control over your data and privacy.
- **Pure text with added Markdown support.** Say goodbye to the overwhelming mental burden of rich formatting and embrace a minimalist approach.
- **Customize and share your notes effortlessly**. With our intuitive sharing features, you can easily collaborate and distribute your notes with others.
- **RESTful API for third-party services.** Embrace the power of integration and unleash new possibilities with our RESTful API support.

## Deploy with Docker in seconds

```bash
docker run -d --name memos -p 5230:5230 -v ~/.memos/:/var/opt/memos ghcr.io/usememos/memos:latest
```

> The `~/.memos/` directory will be used as the data directory on your local machine, while `/var/opt/memos` is the directory of the volume in Docker and should not be modified.

Learn more about [other installation methods](https://www.usememos.com/docs/install).

## Contribution

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. We greatly appreciate any contributions you make. Thank you for being a part of our community! 🥰

<a href="https://github.com/usememos/memos/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=usememos/memos" />
</a>

---

- [Moe Memos](https://memos.moe/) - Third party client for iOS and Android
- [lmm214/memos-bber](https://github.com/lmm214/memos-bber) - Chrome extension
- [Rabithua/memos_wmp](https://github.com/Rabithua/memos_wmp) - WeChat MiniProgram
- [qazxcdswe123/telegramMemoBot](https://github.com/qazxcdswe123/telegramMemoBot) - Telegram bot
- [eallion/memos.top](https://github.com/eallion/memos.top) - Static page rendered with the Memos API
- [eindex/logseq-memos-sync](https://github.com/EINDEX/logseq-memos-sync) - Logseq plugin
- [JakeLaoyu/memos-import-from-flomo](https://github.com/JakeLaoyu/memos-import-from-flomo) - Import data. Support from flomo, wechat reading
- [Quick Memo](https://www.icloud.com/shortcuts/1eaef307112843ed9f91d256f5ee7ad9) - Shortcuts (iOS, iPadOS or macOS)
- [Memos Raycast Extension](https://www.raycast.com/JakeYu/memos) - Raycast extension
- [Memos Desktop](https://github.com/xudaolong/memos-desktop) - Third party client for MacOS and Windows
- [MemosGallery](https://github.com/BarryYangi/MemosGallery) - A static Gallery rendered with the Memos API

## Star history

[![Star History Chart](https://api.star-history.com/svg?repos=usememos/memos&type=Date)](https://star-history.com/#usememos/memos&Date)
