package components

import (
	"encoding/gob"
	"encoding/json"
	"strconv"

	"github.com/SHP-Association/E-learningWeb/backend/pkg/msg"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/ui"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func init() {
	gob.Register(Alert{})
}

type Alert struct {
	ID       string            `json:"id"`
	Type     string            `json:"type"`
	Title    string            `json:"title"`
	Message  string            `json:"message"`
	Fields   map[string]string `json:"fields"`
	Duration int               `json:"duration"`
}

func FlashMessages(r *ui.Request) Node {
	var g Group

	for _, typ := range []msg.Type{
		msg.TypeSuccess,
		msg.TypeInfo,
		msg.TypeWarning,
		msg.TypeError,
	} {
		for _, m := range msg.Get(r.Context, typ) {
			g = append(g, RenderAlert(Alert{
				Type:    string(typ),
				Title:   defaultTitle(typ),
				Message: m,
			}))
		}
	}

	return g
}

func RenderAlert(a Alert) Node {
	fields, _ := json.Marshal(a.Fields)
	duration := a.Duration
	if duration == 0 {
		switch a.Type {
		case "error":
			duration = 8000
		case "warning":
			duration = 6000
		default:
			duration = 4000
		}
	}
	return Div(
		Attr("data-server-alert", ""),
		Attr("data-type", a.Type),
		Attr("data-title", a.Title),
		Attr("data-message", a.Message),
		Attr("data-fields", string(fields)),
		Attr("data-duration", strconv.Itoa(duration)),
		Style("display:none"),
	)
}

func defaultTitle(typ msg.Type) string {
	switch typ {
	case msg.TypeSuccess:
		return "Success"
	case msg.TypeError:
		return "Error"
	case msg.TypeWarning:
		return "Warning"
	case msg.TypeInfo:
		return "Info"
	}
	return "Notification"
}

func AlertContainer() Node {
	return Div(
		ID("alert-container"),
		Style("position:fixed; top:24px; right:24px; z-index:999999; width:400px; max-width:calc(100vw - 48px); display:flex; flex-direction:column; gap:12px; pointer-events:none; isolation:isolate;"),
	)
}

func AlertJS() Node {
	return Script(Raw(`(function() {
  if (window.__alertsInitialized) return;
  window.__alertsInitialized = true;

  const queue = [];
  let active = null;
  const seenMessages = new Set();

  window.showAlert = function(alert) {
    const key = alert.type + ":" + alert.message;
    if (seenMessages.has(key)) return;
    seenMessages.add(key);
    setTimeout(() => seenMessages.delete(key), 1000);

    queue.push(alert);
    if (!active) processQueue();
  };

  function processQueue() {
    if (queue.length === 0) { active = null; return; }
    active = queue.shift();
    const el = createAlertEl(active);
    const container = document.getElementById('alert-container');
    if (container) container.appendChild(el);

    const timer = setTimeout(() => dismiss(el), active.duration || 5000);
    el._timer = timer;
  }

  function dismiss(el) {
    clearTimeout(el._timer);
    el.style.opacity = '0';
    el.style.transform = 'translateX(20px) scale(0.95)';
    setTimeout(() => {
      el.remove();
      processQueue();
    }, 300);
  }

  function createAlertEl(alert) {
    const configs = {
      error:   { accent: 'var(--color-danger)', icon: '✕' },
      warning: { accent: 'var(--color-warning)', icon: '⚠️' },
      success: { accent: 'var(--color-accent)', icon: '✓' },
      info:    { accent: 'var(--color-info)', icon: 'ℹ' },
    };
    const c = configs[alert.type] || configs.info;
    const isLight = document.documentElement.getAttribute('data-theme') === 'light';
    const el = document.createElement('div');
    el.style.cssText = 'background: ' + (isLight ? 'var(--color-card-bg)' : 'color-mix(in srgb, var(--color-card-bg) 85%, transparent)') + '; ' +
      'border: 1.5px solid var(--color-card-border); border-radius: 20px; ' +
      'padding: 20px; pointer-events: all; cursor: default; ' +
      (isLight ? '' : 'backdrop-filter: blur(20px); -webkit-backdrop-filter: blur(20px); ') +
      'box-shadow: var(--shadow-lg); ' +
      'transition: all 0.5s cubic-bezier(0.19, 1, 0.22, 1); ' +
      'opacity: 0; transform: translateX(80px); position: relative; overflow: hidden; z-index: 999999;';

    let fieldsHTML = '';
    if (alert.fields && Object.keys(alert.fields).length > 0) {
      fieldsHTML = '<ul style="margin:12px 0 0 0; padding-left:0; list-style:none; font-size:11px; opacity:0.8;">';
      for (const [field, msg] of Object.entries(alert.fields)) {
        fieldsHTML += '<li style="color:' + c.accent + '; margin-bottom:4px; display:flex; gap:8px;"><strong style="text-transform:uppercase; letter-spacing:0.1em; opacity:0.6;">' + field + ':</strong> <span>' + msg + '</span></li>';
      }
      fieldsHTML += '</ul>';
    }

    el.innerHTML = '<div style="display:flex; justify-content:space-between; align-items:flex-start; gap:20px;">' +
        '<div style="display:flex; gap:16px; align-items:center; flex:1;">' +
          '<div style="width:40px; height:40px; background:' + c.accent + '15; border: 1px solid ' + c.accent + '30; border-radius:14px; display:flex; align-items:center; justify-content:center; flex-shrink:0; font-size:18px; color:' + c.accent + '; font-weight:bold;">' + c.icon + '</div>' +
          '<div style="flex:1;">' +
            '<p style="margin:0; font-family:\'Outfit\', sans-serif; font-size:11px; font-weight:900; color:' + c.accent + '; text-transform:uppercase; letter-spacing:0.15em;">' + alert.title + '</p>' +
            '<p style="margin:2px 0 0; font-family:\'Inter\', sans-serif; font-size:14px; font-weight:500; color:var(--color-primary-text); line-height:1.4;">' + alert.message + '</p>' +
            fieldsHTML +
          '</div>' +
        '</div>' +
        '<button style="background:rgba(255,255,255,0.03); border:1px solid rgba(255,255,255,0.05); border-radius:10px; cursor:pointer; width:28px; height:28px; color:var(--color-secondary-text); display:flex; align-items:center; justify-content:center; flex-shrink:0; transition:all 0.2s; font-size:16px;" onmouseover="this.style.background=\'rgba(255,50,50,0.1)\'; this.style.color=\'#ff4444\'" onmouseout="this.style.background=\'rgba(255,255,255,0.03)\'; this.style.color=\'var(--color-secondary-text)\'">×</button>' +
      '</div>' +
      '<div style="position:absolute; bottom:0; left:0; right:0; height:2px; background:rgba(255,255,255,0.05);">' +
        '<div class="alert-progress" style="height:100%; background:' + c.accent + '; width:100%; transition:width ' + (alert.duration || 5000) + 'ms linear;"></div>' +
      '</div>';

    el.setAttribute('data-alert', '');
    const closeBtn = el.querySelector('button');
    if (closeBtn) closeBtn.addEventListener('click', () => dismiss(el));

    requestAnimationFrame(() => {
      requestAnimationFrame(() => {
        el.style.opacity = '1';
        el.style.transform = 'translateX(0) scale(1)';
        const progress = el.querySelector('.alert-progress');
        if (progress) progress.style.width = '0%';
      });
    });

    return el;
  }

  function processServerAlerts() {
    document.querySelectorAll("[data-server-alert]").forEach(function(el) {
      window.showAlert({
        type:     el.dataset.type,
        title:    el.dataset.title,
        message:  el.dataset.message,
        fields:   JSON.parse(el.dataset.fields || '{}'),
        duration: parseInt(el.dataset.duration || '5000'),
      });
      el.remove();
    });
  }

  processServerAlerts();
  document.addEventListener('DOMContentLoaded', processServerAlerts);
  document.addEventListener('htmx:afterSwap', processServerAlerts);
})();`))
}
