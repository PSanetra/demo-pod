import { environment } from '../../environments/environment';

const baseElement = document.querySelector('base');

export function apiUri(): string {
  const apiUri = environment.apiUri;

  if (/^http(s)?:\/\//.test(apiUri)) {
    return apiUri;
  }

  if (!baseElement) {
    return apiUri;
  }

  const baseHref = baseElement.href || '';

  if (baseHref.endsWith('/') && apiUri.startsWith('/')) {
    return baseHref + apiUri.substr(1);
  }

  if (baseHref.endsWith('/') !== apiUri.startsWith('/')) {
    return baseHref + apiUri;
  }

  return `${baseHref}/${apiUri}`;
}
