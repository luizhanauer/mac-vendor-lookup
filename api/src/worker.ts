export interface VendorInfo {
  macPrefix: string;
  companyName: string;
  maskLength: number;
}

export interface ShardAggregate {
  baseVendor?: string;
  blocks?: VendorInfo[];
}

export interface ApiResponse {
  mac: string;
  vendor: string;
  exactMatch: boolean;
}

export default {
  async fetch(request: Request): Promise<Response> {
    const url = new URL(request.url);
    const rawMac = url.pathname.replace('/', '').trim();

    if (!rawMac) {
      return createErrorResponse(400, "Forneca um MAC address na URL. Ex: /a8:23:fe:08:37:ba");
    }

    const cleanMac = rawMac.replace(/[^a-fA-F0-9]/g, '').toUpperCase();
    
    if (cleanMac.length < 6) {
      return createErrorResponse(400, "MAC address invalido. Minimo de 6 caracteres hexadecimais.");
    }

    const oui = cleanMac.substring(0, 6);
    const p1 = oui.substring(0, 2);
    const p2 = oui.substring(2, 4);
    const p3 = oui.substring(4, 6);

    // Apontando para o seu GitHub Pages real
    const ghUrl = `https://luizhanauer.github.io/mac-vendor-lookup/v1/${p1}/${p2}/${p3}.json`;
    const ghResponse = await fetch(ghUrl);

    if (!ghResponse.ok) {
      return createErrorResponse(404, "Fabricante nao encontrado para este MAC.");
    }

    const shard = await ghResponse.json<ShardAggregate>();
    const result = findExactVendor(cleanMac, shard);

    if (!result) {
       return createErrorResponse(404, "Fabricante nao encontrado nos registros.");
    }

    return new Response(JSON.stringify(result), {
      status: 200,
      headers: {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*'
      }
    });
  }
};

function findExactVendor(cleanMac: string, shard: ShardAggregate): ApiResponse | null {
  const specificBlock = findSpecificBlock(cleanMac, shard.blocks || []);
  
  if (specificBlock) {
    return {
      mac: cleanMac,
      vendor: specificBlock.companyName,
      exactMatch: true
    };
  }

  if (shard.baseVendor) {
    return {
      mac: cleanMac,
      vendor: shard.baseVendor,
      exactMatch: false 
    };
  }

  return null;
}

function findSpecificBlock(cleanMac: string, blocks: VendorInfo[]): VendorInfo | null {
  for (const block of blocks) {
    if (cleanMac.startsWith(block.macPrefix)) {
      return block;
    }
  }
  return null;
}

function createErrorResponse(status: number, message: string): Response {
  return new Response(JSON.stringify({ error: message }), {
    status,
    headers: { 
      'Content-Type': 'application/json',
      'Access-Control-Allow-Origin': '*'
    }
  });
}