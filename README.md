# MAC Vendor Lookup API 🌐

Uma API rápida, gratuita e *serverless* para consultar fabricantes a partir de endereços MAC (Media Access Control). 

🔗 **Links Rápidos:**
* **Repositório:** [github.com/luizhanauer/mac-vendor-lookup](https://github.com/luizhanauer/mac-vendor-lookup)
* **API de Produção:** [mac-vendor-lookup.hanauerlabs.com.br](https://mac-vendor-lookup.hanauerlabs.com.br/)
* **Web App (Em Breve):** [luizhanauer.github.io/mac-vendor-lookup/](https://luizhanauer.github.io/mac-vendor-lookup/)

Este projeto utiliza uma arquitetura híbrida de **Geração Estática (Go)** e **Edge Computing (Cloudflare Workers)** para fornecer respostas em milissegundos com custo de infraestrutura zero. O banco de dados é gerado a partir da base oficial do Wireshark (`manuf`).

---

## 🏗 Arquitetura do Sistema

O sistema foi desenhado visando alta performance, isolamento de domínio e baixo custo. Ele é dividido nas seguintes camadas:

1. **The Generator (Go):** Um script em Go otimizado processa o arquivo bruto `manuf` do Wireshark. Ele aplica o padrão de *Aggregate Root* agrupando os prefixos de 24-bits (OUI) e sub-blocos (como IABs) em milhares de pequenos arquivos JSON fragmentados (Sharding).
2. **Static Data Layer (GitHub Pages - branch `gh-pages`):** Os arquivos JSON fragmentados são hospedados no GitHub Pages. Isso atua como um banco de dados estático e altamente escalável, distribuído globalmente via CDN.
3. **The Edge API (TypeScript):** Um Cloudflare Worker atua como o *API Gateway*. Ele recebe a requisição do usuário, limpa o MAC address, busca o fragmento exato no GitHub Pages (ex: `/v1/A8/23/FE.json`) e processa a lógica de negócio para retornar o fabricante exato, mesmo para sub-redes com máscaras não padrão.
4. **Web Interface (Vue.js - *Em Desenvolvimento*):** Um frontend reativo em Vue.js que será servido nativamente pelo GitHub Pages para fornecer uma interface gráfica intuitiva para consultas manuais, consumindo diretamente a Edge API.

---

## 🚀 Como Usar a API

A API é pública e pode ser consumida via requisições HTTP `GET`.

**Endpoint:**
```http
GET https://mac-vendor-lookup.hanauerlabs.com.br/{mac-address}
```

*(Nota: O MAC address pode ser enviado em qualquer formato: com dois pontos, traços ou apenas os caracteres hexadecimais).*

### Exemplo 1: Match de OUI Padrão (/24)

```bash
curl https://mac-vendor-lookup.hanauerlabs.com.br/a8:23:fe:08:37:ba

```

**Resposta (200 OK):**

```json
{
  "mac": "A823FE0837BA",
  "vendor": "Apple, Inc.",
  "exactMatch": false
}

```

### Exemplo 2: Match Exato em Sub-blocos (IAB /36)

```bash
curl https://mac-vendor-lookup.hanauerlabs.com.br/00:50:C2:4B:41:00

```

**Resposta (200 OK):**

```json
{
  "mac": "0050C24B4100",
  "vendor": "Some Specific Company Ltd.",
  "exactMatch": true
}

```

### Exemplo 3: Erro de Validação

```bash
curl https://mac-vendor-lookup.hanauerlabs.com.br/invalid-mac

```

**Resposta (400 Bad Request):**

```json
{
  "error": "MAC address invalido. Minimo de 6 caracteres hexadecimais."
}

```

---

## 💻 Estrutura do Repositório (Mono-repo)

O projeto segue os princípios de Clean Architecture e separação de responsabilidades:

```text
.
├── api/                        # Cloudflare Worker (TypeScript Strict)
│   ├── src/
│   │   └── worker.ts           # Entry point e lógica de negócio da borda
│   ├── package.json
│   ├── tsconfig.json
│   └── wrangler.toml           # Configuração de deploy da Cloudflare
├── cmd/
│   └── generator/
│       └── main.go             # Entry point do gerador de dados estáticos
├── internal/
│   ├── domain/                 # Entidades e Value Objects (OUI, VendorInfo)
│   └── infra/                  # Parser e Storage (I/O)
├── public/                     # Diretório de saída gerado automaticamente (Ignorado no Git principal)
└── .github/workflows/          # Automações de CI/CD para Go e TS

```

---

## 🛠 Desenvolvimento Local

### 1. Gerando a Base Estática (Go)

Para baixar a base mais recente do Wireshark e gerar os JSONs fragmentados localmente:

```bash
# Baixe a base bruta
curl -L https://www.wireshark.org/download/automated/data/manuf -o manuf.txt

# Execute o gerador
go run cmd/generator/main.go -input manuf.txt -output public
```

### 2. Rodando a API Localmente (TypeScript)

Para testar o Cloudflare Worker em seu ambiente local:

```bash
cd api
npm install
npx wrangler dev
```

*(O Worker local continuará apontando para a sua base estática hospedada no GitHub Pages, conforme configurado em `worker.ts`).*

---

## 🔄 Automação e Deploy (CI/CD)

Este repositório utiliza GitHub Actions para manter a API sempre atualizada sem intervenção manual:

* **Deploy da API:** Qualquer alteração na pasta `api/` aciona o deploy automático para a Cloudflare via Wrangler.
* **Data Updater:** Um cron job semanal fará o download da base atualizada do Wireshark, rodará o gerador em Go e fará o push dos novos JSONs estáticos para o GitHub Pages.

---

## 🤝 Contribuição

Contribuições são bem-vindas! Se você encontrar algum problema ou tiver sugestões para melhorar a aplicação, sinta-se à vontade para abrir uma issue ou enviar um pull request.

Se você gostou do meu trabalho e quer me agradecer, você pode me pagar um café :)

<a href="https://www.paypal.com/donate/?hosted_button_id=SFR785YEYHC4E" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 40px !important;width: 150px !important;" ></a>

## 📄 Licença

Este projeto está licenciado sob a Licença MIT. Consulte o arquivo `LICENSE` para obter mais informações.