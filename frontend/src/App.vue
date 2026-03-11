<script setup lang="ts">
import { ref, computed } from 'vue';

interface ApiResponse {
  mac: string;
  vendor: string;
  exactMatch: boolean;
}

const macInput = ref<string>('');
const isLoading = ref<boolean>(false);
const errorMsg = ref<string>('');
const result = ref<ApiResponse | null>(null);
const isCopied = ref<boolean>(false);

const rawMacClean = computed((): string => {
  return macInput.value.replace(/[^a-fA-F0-9]/g, '').toLowerCase();
});

const curlCommand = computed((): string => {
  if (!rawMacClean.value) return '';
  return `curl -s https://mac-vendor-lookup.hanauerlabs.com.br/${rawMacClean.value}`;
});

const handleLookup = async (): Promise<void> => {
  if (!macInput.value) {
    errorMsg.value = 'Por favor, forneça um endereço MAC para consultar.';
    return;
  }

  if (rawMacClean.value.length < 6) {
    errorMsg.value = 'O MAC Address deve conter pelo menos 6 caracteres hexadecimais válidos.';
    return;
  }

  isLoading.value = true;
  errorMsg.value = '';
  result.value = null;
  isCopied.value = false;

  try {
    const response = await fetch(`https://mac-vendor-lookup.hanauerlabs.com.br/${rawMacClean.value}`);
    
    if (!response.ok) {
      const errData = await response.json();
      throw new Error(errData.error || 'Fabricante não encontrado na base de dados.');
    }

    result.value = await response.json() as ApiResponse;
  } catch (err) {
    errorMsg.value = err instanceof Error ? err.message : 'Falha na comunicação com a API. Tente novamente.';
  } finally {
    isLoading.value = false;
  }
};

const copyToClipboard = async (text: string): Promise<void> => {
  try {
    await navigator.clipboard.writeText(text);
    isCopied.value = true;
    
    setTimeout(() => {
      isCopied.value = false;
    }, 2500);
  } catch (err) {
    console.error('Falha ao copiar para a área de transferência', err);
  }
};
</script>

<template>
  <div class="min-h-screen relative flex flex-col justify-between selection:bg-[#289f91] selection:text-white">
    
    <div class="fixed inset-0 z-0 pointer-events-none overflow-hidden flex items-center justify-center">
      <svg id="nic-circuit-bg" class="w-full h-full object-cover opacity-60" preserveAspectRatio="xMidYMid slice" viewBox="0 0 1280 768" xmlns="http://www.w3.org/2000/svg">
        <rect width="100%" height="100%" fill="#0a0b10"/>
        
        <g class="circuit-effect">
          <g stroke="#289f91" stroke-width="2" opacity="0.25" fill="none">
            <rect x="-20" y="220" width="140" height="240" rx="5" />
            <path d="M 100 240 L 100 440" stroke-dasharray="4 4" />
            
            <rect x="220" y="260" width="80" height="160" rx="3" />
            
            <rect x="550" y="200" width="240" height="240" rx="8" />
            <rect x="590" y="240" width="160" height="160" rx="4" opacity="0.5" />
            <path d="M 540 240 L 550 240 M 540 260 L 550 260 M 540 280 L 550 280 M 540 300 L 550 300 M 540 320 L 550 320 M 540 340 L 550 340 M 540 360 L 550 360 M 540 380 L 550 380 M 540 400 L 550 400" />
            <path d="M 600 440 L 600 450 M 630 440 L 630 450 M 660 440 L 660 450 M 690 440 L 690 450 M 720 440 L 720 450" />
            <path d="M 790 260 L 800 260 M 790 300 L 800 300 M 790 340 L 800 340 M 790 380 L 800 380" />

            <rect x="900" y="220" width="70" height="90" rx="2" />
            <path d="M 885 240 L 900 240 M 885 260 L 900 260 M 885 280 L 900 280 M 885 300 L 900 300" />
            <path d="M 970 240 L 985 240 M 970 260 L 985 260 M 970 280 L 985 280 M 970 300 L 985 300" />

            <rect x="420" y="120" width="60" height="40" rx="10" />
            
            <path d="M 400 768 L 400 710 L 440 710 L 440 730 L 460 730 L 460 710 L 950 710 L 950 768" />
          </g>

          <g>
            <rect class="pcie-pad" x="480" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="500" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="520" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="540" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="560" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="580" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="600" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="620" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="640" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="660" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="680" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="700" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="720" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="740" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="760" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="780" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="800" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="820" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="840" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="860" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="880" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="900" y="710" width="12" height="58" />
            <rect class="pcie-pad" x="920" y="710" width="12" height="58" />
          </g>

          <g stroke="#289f91" stroke-width="1.5" opacity="0.15" fill="none">
            <path d="M 120 280 L 220 280" />
            <path d="M 120 300 L 220 300" />
            <path d="M 120 360 L 220 360" />
            <path d="M 120 380 L 220 380" />
            <path d="M 300 280 L 380 280 L 440 240 L 550 240" />
            <path d="M 300 300 L 390 300 L 440 260 L 550 260" />
            <path d="M 300 360 L 420 360 L 460 320 L 550 320" />
            <path d="M 300 380 L 410 380 L 450 340 L 550 340" />
            <path d="M 300 410 L 380 410 L 430 380 L 550 380" />
            <path d="M 600 440 L 600 600 L 520 680 L 520 710" />
            <path d="M 630 440 L 630 580 L 560 650 L 560 710" />
            <path d="M 660 440 L 660 620 L 640 640 L 640 710" />
            <path d="M 690 440 L 690 620 L 720 650 L 720 710" />
            <path d="M 720 440 L 720 580 L 800 660 L 800 710" />
            <path d="M 790 260 L 850 260 L 870 240 L 900 240" />
            <path d="M 790 300 L 840 300 L 860 280 L 900 280" />
            <path d="M 450 160 L 450 180 L 500 230 L 550 230" />
          </g>

          <path class="circuit-path" style="animation-duration: 3s; animation-delay: 0s;" stroke="#289f91" d="M 550 240 L 440 240 L 380 280 L 300 280" />
          <path class="circuit-path" style="animation-duration: 3s; animation-delay: 0.2s;" stroke="#289f91" d="M 300 280 L 220 280 M 120 280 L -20 280" />
          <path class="circuit-path" style="animation-duration: 2.5s; animation-delay: 1.5s;" stroke="#289f91" d="M -20 360 L 120 360 M 220 360 L 300 360" />
          <path class="circuit-path" style="animation-duration: 2.5s; animation-delay: 1.7s;" stroke="#289f91" d="M 300 360 L 420 360 L 460 320 L 550 320" />
          <path class="circuit-path" style="animation-duration: 2.8s; animation-delay: 0.8s;" stroke="#289f91" d="M -20 380 L 120 380 M 220 380 L 300 380 L 410 380 L 450 340 L 550 340" />
          <path class="circuit-path" style="animation-duration: 1.8s; animation-delay: 0.5s;" stroke="#289f91" d="M 520 710 L 520 680 L 600 600 L 600 440" />
          <path class="circuit-path" style="animation-duration: 2.2s; animation-delay: 1.2s;" stroke="#289f91" d="M 660 440 L 660 620 L 640 640 L 640 710" />
          <path class="circuit-path" style="animation-duration: 1.5s; animation-delay: 2.1s;" stroke="#289f91" d="M 800 710 L 800 660 L 720 580 L 720 440" />
          <path class="circuit-path" style="animation-duration: 4s; animation-delay: 3s;" stroke="#289f91" stroke-dasharray="30 1000" d="M 900 240 L 870 240 L 850 260 L 790 260" />
          <path class="circuit-path" style="animation-duration: 4s; animation-delay: 3.2s;" stroke="#289f91" stroke-dasharray="30 1000" d="M 790 300 L 840 300 L 860 280 L 900 280" />
          <path class="circuit-path" style="animation-duration: 1s; animation-delay: 0s;" stroke="#289f91" stroke-dasharray="10 1000" d="M 450 160 L 450 180 L 500 230 L 550 230" />

          <g fill="#289f91" opacity="0.4">
            <circle cx="300" cy="280" r="3" />
            <circle cx="300" cy="300" r="3" />
            <circle cx="300" cy="360" r="3" />
            <circle cx="300" cy="380" r="3" />
            <circle cx="300" cy="410" r="3" />
            <circle cx="450" cy="160" r="2.5" />
            <circle cx="380" cy="280" r="2" />
            <circle cx="390" cy="300" r="2" />
            <circle cx="420" cy="360" r="2" />
            <circle cx="410" cy="380" r="2" />
            <circle cx="380" cy="410" r="2" />
            <circle cx="520" cy="680" r="2.5" />
            <circle cx="560" cy="650" r="2.5" />
            <circle cx="640" cy="640" r="2.5" />
            <circle cx="720" cy="650" r="2.5" />
            <circle cx="800" cy="660" r="2.5" />
            <circle cx="850" cy="260" r="2" />
            <circle cx="840" cy="300" r="2" />
          </g>
        </g>
      </svg>
    </div>
    
    <main class="flex-grow flex flex-col items-center justify-center p-6 z-10 w-full max-w-3xl mx-auto">
      
      <header class="text-center mb-10">
        <h1 class="text-4xl font-bold text-white mb-4">MAC Vendor Lookup</h1>
        <p class="text-gray-400 text-sm md:text-base max-w-xl mx-auto leading-relaxed">
          Consulta instantânea de fabricantes a partir de endereços MAC. Nossa API resolve identificadores OUI e IAB processados diretamente da base oficial do Wireshark.
        </p>
      </header>

      <div class="w-full bg-[#12141c]/90 border border-gray-800 backdrop-blur-md p-8 rounded-2xl shadow-2xl">
        
        <form @submit.prevent="handleLookup" class="flex flex-col sm:flex-row gap-3 mb-6">
          <input 
            v-model="macInput" 
            type="text" 
            placeholder="Ex: A8:23:FE:08:37:BA" 
            class="font-code flex-grow bg-[#0a0b10] border border-gray-700 text-[#289f91] px-5 py-4 rounded-xl outline-none focus:border-[#289f91] focus:ring-1 focus:ring-[#289f91] transition-all placeholder-gray-600 text-lg uppercase"
            :disabled="isLoading"
          />
          <button 
            type="submit" 
            class="bg-[#289f91] hover:bg-[#208276] text-white font-semibold px-8 py-4 rounded-xl transition-all disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center min-w-[160px] shadow-lg shadow-[#289f91]/20"
            :disabled="isLoading"
          >
            <span v-if="!isLoading">Consultar</span>
            <span v-if="isLoading" class="flex items-center gap-2">
              <svg class="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
              Buscando
            </span>
          </button>
        </form>

        <div v-if="errorMsg" class="bg-red-500/10 border border-red-500/30 text-red-400 px-5 py-4 rounded-xl mb-6 text-sm flex items-center gap-3">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="8" x2="12" y2="12"></line><line x1="12" y1="16" x2="12.01" y2="16"></line></svg>
          {{ errorMsg }}
        </div>

        <div v-if="result" class="space-y-6 animate-fade-in mt-8">
          
          <div class="p-6 border border-gray-700 rounded-xl bg-[#0a0b10] flex flex-col sm:flex-row items-start sm:items-center justify-between gap-4">
            <div>
              <span class="block text-gray-500 text-xs uppercase tracking-wider mb-1 font-semibold">Fabricante Identificado</span>
              <span class="text-2xl font-bold text-white">{{ result.vendor }}</span>
            </div>
            <div v-if="result.exactMatch" class="text-xs bg-[#289f91]/20 text-[#289f91] px-3 py-1.5 rounded-full border border-[#289f91]/30 font-medium">
              Correspondência Exata (IAB)
            </div>
          </div>

          <div class="space-y-2">
            <span class="text-xs text-gray-400 font-medium ml-1 uppercase tracking-wider">Resposta da API (JSON)</span>
            <div class="relative group">
              <pre class="font-code bg-[#0a0b10] border border-gray-800 rounded-xl p-5 text-sm text-[#289f91] overflow-x-auto"><code>{{ JSON.stringify(result, null, 2) }}</code></pre>
            </div>
          </div>

          <div class="space-y-2">
            <span class="text-xs text-gray-400 font-medium ml-1 uppercase tracking-wider">Comando cURL para Terminal</span>
            <div class="relative flex items-center bg-[#0a0b10] border border-gray-800 rounded-xl p-1.5 transition-colors focus-within:border-gray-600">
              <input 
                type="text" 
                readonly 
                :value="curlCommand"
                class="font-code bg-transparent text-gray-300 text-sm w-full p-3 outline-none"
              />
              <button 
                @click="copyToClipboard(curlCommand)" 
                class="ml-2 px-5 py-2.5 rounded-lg text-sm font-semibold transition-all flex items-center gap-2"
                :class="isCopied 
                  ? 'bg-green-600 hover:bg-green-500 text-white shadow-lg shadow-green-900/50' 
                  : 'bg-gray-800 hover:bg-gray-700 text-gray-200'"
              >
                <svg v-if="!isCopied" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
                {{ isCopied ? 'Copiado!' : 'Copiar' }}
              </button>
            </div>
          </div>

        </div>
      </div>
    </main>

    <footer class="z-10 py-8 border-t border-gray-800/50 bg-[#0a0b10]/80 backdrop-blur-md mt-10">
      <div class="container mx-auto flex flex-col items-center justify-center gap-5">
        <a 
          href="https://github.com/luizhanauer/mac-vendor-lookup" 
          target="_blank" 
          rel="noopener noreferrer"
          class="text-gray-400 hover:text-white transition-colors flex items-center gap-2 text-sm font-medium"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="currentColor"><path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/></svg>
          luizhanauer/mac-vendor-lookup
        </a>
        
        <a 
          href="https://www.paypal.com/donate/?hosted_button_id=SFR785YEYHC4E" 
          target="_blank" 
          rel="noopener noreferrer"
          class="hover:scale-105 transition-transform origin-center"
        >
          <img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" class="h-9" />
        </a>
      </div>
    </footer>

  </div>
</template>

<style scoped>
.animate-fade-in {
  animation: fadeIn 0.4s ease-out forwards;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

@keyframes circuit-run {
  0% { stroke-dashoffset: 1000; opacity: 0; }
  10% { opacity: 0.8; }
  90% { opacity: 0.8; }
  100% { stroke-dashoffset: 0; opacity: 0; }
}
.circuit-path {
  fill: none;
  stroke-width: 2.5;
  stroke-linecap: round;
  stroke-linejoin: round;
  stroke-dasharray: 80 1000;
  animation: circuit-run linear infinite;
}
.pcie-pad {
  fill: #289f91;
  opacity: 0.2;
}
</style>