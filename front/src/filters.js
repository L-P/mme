export default {
  maybeHex(v, width) {
    if (typeof v === 'number') {
      const hex = v.toString(16).toUpperCase().padStart(width, '0');
      return `0x${hex}`;
    }

    return v;
  },

  hex(v, width) {
    const hex = v.toString(16).toUpperCase().padStart(width, '0');
    return `0x${hex}`;
  },

  bool(v) {
    return v ? 't' : 'f';
  },

  humanizeBytes(v) {
    if (v === 0) {
      return '0 B';
    }

    const units = ['B', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'ZiB', 'YiB'];
    const i = Math.min(units.length - 1, Math.floor(Math.log(v) / Math.log(1024)));
    let rounded = v / (1024 ** i);
    rounded = +rounded.toFixed(2);


    return `${rounded} ${units[i]}`;
  },
};
