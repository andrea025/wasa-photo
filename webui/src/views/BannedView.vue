<script>
export default {
  name: 'BannedView',
  data: function() {
    return {
      errormsg: null,
      banned: [],
    };
  },
  methods: {
    async getBanned() {
      this.errormsg = null;
      try {
        const response = await this.$axios.get('/users/' + this.$route.params.id + '/banned');
        this.banned = response.data.data;
      } catch (e) {
        switch (e.response.status) {
          case 400:
            this.errormsg = 'Ops, there was something wrong with your request.';
            break;
          case 401:
            this.errormsg = 'You need to login in order to perform this action.';
            break;
          case 403:
            this.errormsg = 'Action forbidden: you cannot see the list of banned users from another user.';
            break;
          case 404:
            this.errormsg = 'Ops, user not found.';
            break;
          case 500:
            this.errormsg = 'Ops, there was an internal problem with the server.';
            break;
          default:
            this.errormsg = e.toString();
        }
      }
    },
  },
  mounted() {
    this.getBanned();
  },
};
</script>


<template>
<div class="section page mt-60" v-if="this.errormsg">
    <ErrorMsg :msg="this.errormsg" />
</div>
  <div style="--fds-black:#000000; --fds-black-alpha-05:rgba(0, 0, 0, 0.05); --fds-black-alpha-10:rgba(0, 0, 0, 0.1); --fds-black-alpha-15:rgba(0, 0, 0, 0.15); --fds-black-alpha-20:rgba(0, 0, 0, 0.2); --fds-black-alpha-30:rgba(0, 0, 0, 0.3); --fds-black-alpha-40:rgba(0, 0, 0, 0.4); --fds-black-alpha-50:rgba(0, 0, 0, 0.5); --fds-black-alpha-60:rgba(0, 0, 0, 0.6); --fds-black-alpha-80:rgba(0, 0, 0, 0.8); --fds-blue-05:#ECF3FF; --fds-blue-30:#AAC9FF; --fds-blue-40:#77A7FF; --fds-blue-60:#1877F2; --fds-blue-70:#2851A3; --fds-blue-80:#1D3C78; --fds-button-text:#444950; --fds-comment-background:#F2F3F5; --fds-dark-mode-gray-35:#CCCCCC; --fds-dark-mode-gray-50:#828282; --fds-dark-mode-gray-70:#4A4A4A; --fds-dark-mode-gray-80:#373737; --fds-dark-mode-gray-90:#282828; --fds-dark-mode-gray-100:#1C1C1C; --fds-gray-00:#F5F6F7; --fds-gray-05:#F2F3F5; --fds-gray-10:#EBEDF0; --fds-gray-20:#DADDE1; --fds-gray-25:#CCD0D5; --fds-gray-30:#BEC3C9; --fds-gray-45:#8D949E; --fds-gray-70:#606770; --fds-gray-80:#444950; --fds-gray-90:#303338; --fds-gray-100:#1C1E21; --fds-green-55:#00A400; --fds-highlight:#3578E5; --fds-highlight-cell-background:#ECF3FF; --fds-primary-icon:#1C1E21; --fds-primary-text:#1C1E21; --fds-red-55:#FA383E; --fds-soft:cubic-bezier(0.08,0.52,0.52,1); --fds-spectrum-aluminum-tint-70:#E4F0F6; --fds-spectrum-blue-gray-tint-70:#CFD1D5; --fds-spectrum-cherry:#F35369; --fds-spectrum-cherry-tint-70:#FBCCD2; --fds-spectrum-grape-tint-70:#DDD5F0; --fds-spectrum-grape-tint-90:#F4F1FA; --fds-spectrum-lemon-dark-1:#F5C33B; --fds-spectrum-lemon-tint-70:#FEF2D1; --fds-spectrum-lime:#A3CE71; --fds-spectrum-lime-tint-70:#E4F0D5; --fds-spectrum-orange-tint-70:#FCDEC5; --fds-spectrum-orange-tint-90:#FEF4EC; --fds-spectrum-seafoam-tint-70:#CAEEF9; --fds-spectrum-slate-dark-2:#89A1AC; --fds-spectrum-slate-tint-70:#EAEFF2; --fds-spectrum-teal:#6BCEBB; --fds-spectrum-teal-dark-1:#4DBBA6; --fds-spectrum-teal-dark-2:#31A38D; --fds-spectrum-teal-tint-70:#D2F0EA; --fds-spectrum-teal-tint-90:#F0FAF8; --fds-spectrum-tomato:#FB724B; --fds-spectrum-tomato-tint-30:#F38E7B; --fds-spectrum-tomato-tint-90:#FDEFED; --fds-strong:cubic-bezier(0.12,0.8,0.32,1); --fds-white:#FFFFFF; --fds-white-alpha-05:rgba(255, 255, 255, 0.05); --fds-white-alpha-10:rgba(255, 255, 255, 0.1); --fds-white-alpha-20:rgba(255, 255, 255, 0.2); --fds-white-alpha-30:rgba(255, 255, 255, 0.3); --fds-white-alpha-40:rgba(255, 255, 255, 0.4); --fds-white-alpha-50:rgba(255, 255, 255, 0.5); --fds-white-alpha-60:rgba(255, 255, 255, 0.6); --fds-white-alpha-80:rgba(255, 255, 255, 0.8); --fds-yellow-20:#FFBA00; --accent:#0095F6; --always-white:#FFFFFF; --always-black:black; --always-dark-gradient:linear-gradient(rgba(0,0,0,0), rgba(0,0,0,0.6)); --always-dark-overlay:rgba(0, 0, 0, 0.4); --always-light-overlay:rgba(255, 255, 255, 0.4); --always-gray-40:#65676B; --always-gray-75:#BCC0C4; --always-gray-95:#F0F2F5; --attachment-footer-background:#F0F2F5; --background-deemphasized:#F0F2F5; --base-blue:#1877F2; --base-cherry:#F3425F; --base-grape:#9360F7; --base-lemon:#F7B928; --base-lime:#45BD62; --base-pink:#FF66BF; --base-seafoam:#54C7EC; --base-teal:#2ABBA7; --base-tomato:#FB724B; --blue-link:#00376B; --border-focused:#65676B; --card-background:#FFFFFF; --card-background-flat:#F7F8FA; --comment-background:#F0F2F5; --comment-footer-background:#F6F9FA; --dataviz-primary-1:rgb(48,200,180); --disabled-button-background:rgba(0, 149, 246, 0.3); --disabled-button-text:#FFFFFF; --disabled-icon:#BCC0C4; --disabled-text:#BCC0C4; --divider:#DBDBDB; --event-date:#F3425F; --fb-wordmark:#1877F2; --filter-accent:invert(39%) sepia(57%) saturate(200%) saturate(200%) saturate(200%) saturate(200%) saturate(200%) saturate(147.75%) hue-rotate(202deg) brightness(97%) contrast(96%); --filter-always-white:invert(100%); --filter-disabled-icon:invert(80%) sepia(6%) saturate(200%) saturate(120%) hue-rotate(173deg) brightness(98%) contrast(89%); --filter-placeholder-icon:invert(59%) sepia(11%) saturate(200%) saturate(135%) hue-rotate(176deg) brightness(96%) contrast(94%); --filter-primary-icon:invert(8%) sepia(10%) saturate(200%) saturate(200%) saturate(166%) hue-rotate(177deg) brightness(104%) contrast(91%); --filter-secondary-icon:invert(39%) sepia(21%) saturate(200%) saturate(109.5%) hue-rotate(174deg) brightness(94%) contrast(86%); --filter-warning-icon:invert(77%) sepia(29%) saturate(200%) saturate(200%) saturate(200%) saturate(200%) saturate(200%) saturate(128%) hue-rotate(359deg) brightness(102%) contrast(107%); --filter-blue-link-icon:invert(30%) sepia(98%) saturate(200%) saturate(200%) saturate(200%) saturate(166.5%) hue-rotate(192deg) brightness(91%) contrast(101%); --filter-positive:invert(37%) sepia(61%) saturate(200%) saturate(200%) saturate(200%) saturate(200%) saturate(115%) hue-rotate(91deg) brightness(97%) contrast(105%); --filter-negative:invert(25%) sepia(33%) saturate(200%) saturate(200%) saturate(200%) saturate(200%) saturate(200%) saturate(200%) saturate(110%) hue-rotate(345deg) brightness(132%) contrast(96%); --glimmer-spinner-icon:#65676B; --hero-banner-background:#FFFFFF; --hosted-view-selected-state:rgba(45, 136, 255, 0.1); --highlight-bg:#E7F3FF; --hover-overlay:rgba(0, 0, 0, 0.05); --list-cell-chevron:#65676B; --media-hover:rgba(68, 73, 80, 0.15); --media-inner-border:rgba(0, 0, 0, 0.1); --media-outer-border:#FFFFFF; --media-pressed:rgba(68, 73, 80, 0.35); --messenger-card-background:#FFFFFF; --messenger-reply-background:#F0F2F5; --overlay-alpha-80:rgba(0, 0, 0, 0.65); --overlay-on-media:rgba(0, 0, 0, 0.6); --nav-bar-background:#FFFFFF; --nav-bar-background-gradient:linear-gradient(to top, #FFFFFF, rgba(255,255,255.9), rgba(255,255,255,0.7), rgba(255,255,255,0.4), rgba(255,255,255,0)); --nav-bar-background-gradient-wash:linear-gradient(to top, #F0F2F5, rgba(240,242,245.9), rgba(240,242,245,0.7), rgba(240,242,245,0.4), rgba(240,242,245,0)); --negative:hsl(350, 87%, 55%); --negative-background:hsl(350, 87%, 55%, 20%); --new-notification-background:#E7F3FF; --non-media-pressed:rgba(68, 73, 80, 0.15); --non-media-pressed-on-dark:rgba(255, 255, 255, 0.3); --notification-badge:#e41e3f; --placeholder-icon:#65676B; --placeholder-image:rgb(164, 167, 171); --placeholder-text:#65676B; --placeholder-text-on-media:rgba(255, 255, 255, 0.5); --popover-background:#FFFFFF; --positive:#31A24C; --positive-background:#DEEFE1; --press-overlay:rgba(0, 0, 0, 0.1); --primary-button-background:#0095F6; --primary-button-icon:#FFFFFF; --primary-button-pressed:#77A7FF; --primary-button-text:#FFFFFF; --primary-deemphasized-button-background:rgba(0, 149, 246, 0.1); --primary-deemphasized-button-pressed:rgba(0, 149, 246, 0.05); --primary-deemphasized-button-pressed-overlay:rgba(0, 149, 246, 0.15); --primary-deemphasized-button-text:#0095F6; --primary-icon:#262626; --primary-text:#262626; --primary-text-on-media:#FFFFFF; --primary-web-focus-indicator:#D24294; --progress-ring-neutral-background:rgba(0, 0, 0, 0.2); --progress-ring-neutral-foreground:#000000; --progress-ring-on-media-background:rgba(255, 255, 255, 0.2); --progress-ring-on-media-foreground:#FFFFFF; --progress-ring-blue-background:rgba(24, 119, 242, 0.2); --progress-ring-blue-foreground:hsl(214, 89%, 52%); --progress-ring-disabled-background:rgba(190,195,201, 0.2); --progress-ring-disabled-foreground:#BEC3C9; --rating-star-active:#EB660D; --scroll-thumb:#BCC0C4; --scroll-shadow:0 1px 2px rgba(0, 0, 0, 0.1), 0 -1px rgba(0, 0, 0, 0.1) inset; --secondary-button-background:transparent; --secondary-button-background-floating:#ffffff; --secondary-button-background-on-dark:rgba(0, 0, 0, 0.4); --secondary-button-pressed:rgba(0, 0, 0, 0.05); --secondary-button-stroke:transparent; --secondary-button-text:#0095F6; --secondary-icon:#8E8E8E; --secondary-text:#8E8E8E; --secondary-text-on-media:rgba(255, 255, 255, 0.9); --section-header-text:#4B4C4F; --shadow-1:rgba(0, 0, 0, 0.1); --shadow-2:rgba(0, 0, 0, 0.2); --shadow-5:rgba(0, 0, 0, 0.5); --shadow-8:rgba(0, 0, 0, 0.8); --shadow-inset:rgba(255, 255, 255, 0.5); --shadow-elevated:0px 5px 12px rgba(52, 72, 84, 0.2); --shadow-persistent:0px 0px 12px rgba(52, 72, 84, 0.05); --shadow-primary:0px 5px 12px rgba(52, 72, 84, 0.2); --surface-background:#FFFFFF; --switch-active:hsl(214, 89%, 52%); --text-highlight:rgba(24, 119, 242, 0.2); --text-input-background:#FFFFFF; --toast-background:#FFFFFF; --toast-text:#1C2B33; --toast-text-link:#216FDB; --toggle-active-background:#E7F3FF; --toggle-active-icon:rgb(24, 119, 242); --toggle-active-text:rgb(24, 119, 242); --toggle-button-active-background:#E7F3FF; --wash:#FAFAFA; --web-wash:#FAFAFA; --warning:hsl(40, 89%, 52%); --fb-logo-color:#2D88FF; --dialog-anchor-vertical-padding:56px; --header-height:0px; --global-panel-width:0px; --global-panel-width-expanded:0px; --alert-banner-corner-radius:8px; --button-corner-radius:4px; --button-corner-radius-medium:10px; --button-corner-radius-large:12px; --button-height-large:40px; --button-height-medium:36px; --button-padding-horizontal-large:16px; --button-padding-horizontal-medium:16px; --button-icon-padding-large:16px; --button-icon-padding-medium:16px; --button-inner-icon-spacing-large:3px; --button-inner-icon-spacing-medium:3px; --blueprint-button-height-medium:40px; --blueprint-button-height-large:48px; --card-corner-radius:4px; --card-box-shadow:0 12px 28px 0 var(--shadow-2), 0 2px 4px 0 var(--shadow-1); --card-padding-vertical:20px; --chip-corner-radius:6px; --dialog-corner-radius:8px; --glimmer-corner-radius:8px; --image-corner-radius:4px; --input-corner-radius:4px; --nav-list-cell-corner-radius:8px; --list-cell-corner-radius:8px; --list-cell-min-height:52px; --list-cell-padding-vertical:20px; --list-cell-padding-vertical-with-addon:14px; --nav-list-cell-min-height:0px; --nav-list-cell-padding-vertical:16px; --nav-list-cell-padding-vertical-with-addon:16px; --text-input-multi-padding-between-text-scrollbar:20px; --text-input-multi-padding-scrollbar:16px; --toast-corner-radius:4px; --text-input-caption-margin-top:10px; --text-input-label-top:22px; --text-input-min-height:64px; --text-input-padding-vertical:12px; --fds-animation-enter-exit-in:cubic-bezier(0.14, 1, 0.34, 1); --fds-animation-enter-exit-out:cubic-bezier(0.45, 0.1, 0.2, 1); --fds-animation-swap-shuffle-in:cubic-bezier(0.14, 1, 0.34, 1); --fds-animation-swap-shuffle-out:cubic-bezier(0.45, 0.1, 0.2, 1); --fds-animation-move-in:cubic-bezier(0.17, 0.17, 0, 1); --fds-animation-move-out:cubic-bezier(0.17, 0.17, 0, 1); --fds-animation-expand-collapse-in:cubic-bezier(0.17, 0.17, 0, 1); --fds-animation-expand-collapse-out:cubic-bezier(0.17, 0.17, 0, 1); --fds-animation-passive-move-in:cubic-bezier(0.5, 0, 0.1, 1); --fds-animation-passive-move-out:cubic-bezier(0.5, 0, 0.1, 1); --fds-animation-quick-move-in:cubic-bezier(0.1, 0.9, 0.2, 1); --fds-animation-quick-move-out:cubic-bezier(0.1, 0.9, 0.2, 1); --fds-animation-fade-in:cubic-bezier(0, 0, 1, 1); --fds-animation-fade-out:cubic-bezier(0, 0, 1, 1); --fds-duration-extra-extra-short-in:100ms; --fds-duration-extra-extra-short-out:100ms; --fds-duration-extra-short-in:200ms; --fds-duration-extra-short-out:150ms; --fds-duration-short-in:280ms; --fds-duration-short-out:200ms; --fds-duration-medium-in:400ms; --fds-duration-medium-out:350ms; --fds-duration-long-in:500ms; --fds-duration-long-out:350ms; --fds-duration-extra-long-in:1000ms; --fds-duration-extra-long-out:1000ms; --fds-duration-none:0ms; --fds-fast:200ms; --fds-slow:400ms; --font-family-apple:system-ui, -apple-system, BlinkMacSystemFont, &quot;.SFNSText-Regular&quot;, sans-serif; --font-family-code:ui-monospace, Menlo, Consolas, Monaco, monospace; --font-family-default:Helvetica, Arial, sans-serif; --font-family-segoe:Segoe UI Historic, Segoe UI, Helvetica, Arial, sans-serif; --body-font-family:Placeholder Font; --body-font-size:0.9375rem; --body-font-weight:400; --body-line-height:1.3333; --body-emphasized-font-family:Placeholder Font; --body-emphasized-font-size:0.9375rem; --body-emphasized-font-weight:600; --body-emphasized-line-height:1.3333; --headline1-font-family:Optimistic Display Bold, system-ui, sans-serif; --headline1-font-size:1.75rem; --headline1-font-weight:700; --headline1-line-height:1.2143; --headline2-font-family:Optimistic Display Bold, system-ui, sans-serif; --headline2-font-size:1.5rem; --headline2-font-weight:700; --headline2-line-height:1.25; --headline3-font-family:Optimistic Display Bold, system-ui, sans-serif; --headline3-font-size:1.0625rem; --headline3-font-weight:700; --headline3-line-height:1.2941; --meta-font-family:Placeholder Font; --meta-font-size:0.8125rem; --meta-font-weight:400; --meta-line-height:1.3846; --meta-emphasized-font-family:Placeholder Font; --meta-emphasized-font-size:0.8125rem; --meta-emphasized-font-weight:600; --meta-emphasized-line-height:1.3846; --primary-label-font-family:Optimistic Display Medium, system-ui, sans-serif; --primary-label-font-size:1.0625rem; --primary-label-font-weight:500; --primary-label-line-height:1.2941; --secondary-label-font-family:Placeholder Font; --secondary-label-font-size:0.9375rem; --secondary-label-font-weight:500; --secondary-label-line-height:1.3333; --text-input-field-font-family:Placeholder Font; --text-input-field-font-size:1rem; --text-input-field-font-weight:500; --text-input-field-line-height:1.2941; --text-input-label-font-family:Placeholder Font; --text-input-label-font-size:17px; --text-input-label-font-size-scale-multiplier:0.75; --text-input-label-font-weight:400; --text-input-label-line-height:1.2941; --dataviz-primary-2:rgb(134,218,255); --dataviz-primary-3:rgb(95,170,255); --dataviz-secondary-1:rgb(118,62,230); --dataviz-secondary-2:rgb(147,96,247); --dataviz-secondary-3:rgb(219,26,139); --dataviz-supplementary-1:rgb(255,122,105); --dataviz-supplementary-2:rgb(241,168,23); --dataviz-supplementary-3:rgb(49,162,76); --dataviz-supplementary-4:rgb(50,52,54); --base-unit:4px; --blue-0:#f5fbff; --blue-2:#b3dbff; --blue-4:#47afff; --blue-5:#0095f6; --blue-6:#0074cc; --blue-7:#0057a3; --blue-8:#00376b; --blue-9:#002952; --breakpoint-medium-width:1536px; --breakpoint-small-width:1024px; --challenge-width:460px; --clr-separator:#efefef; --clr_red_dark_30:#af2634; --cluster-card-border-radius:8px; --creation-composer-height:81px; --creation-header-composer-height:126px; --creation-header-height:43px; --creation-min-padding-x:32px; --creation-modal-max-height:898px; --creation-modal-min-height:391px; --creation-padding-x:64px; --creation-padding-y:112px; --creation-settings-width:340px; --cyan-5:#27c4f5; --desktop-center-feed-min-width-breakpoint:1500px; --desktop-cluster-card-height:89px; --desktop-cluster-card-width:140px; --desktop-collapsed-nav-height:52px; --desktop-grid-item-margin:28px; --desktop-in-feed-story-item-height:208px; --desktop-in-feed-story-item-width:116px; --desktop-large-modal-max-height:781px; --desktop-large-modal-max-width:1491px; --desktop-large-modal-min-height:664px; --desktop-large-modal-min-width:908px; --desktop-nav-anim-duration:0.2s; --desktop-nav-height:60px; --desktop-nav-search-box-large-width:268px; --desktop-nav-search-box-width:215px; --desktop-regular-nav-height:77px; --desktop-skinny-nav-height:60px; --direct-attachment-image-grid-border-size:1px; --direct-attachment-image-grid-item-size:78px; --direct-attachment-image-grid-width:236px; --direct-attachment-story-height:150px; --direct-attachment-story-large-height:256px; --direct-attachment-story-large-width:164px; --direct-attachment-story-width:84px; --direct-message-margin:8px; --direct-message-max-width:236px; --extra-small-screen-max:413px; --fb-connect-blue:#4f67b0; --fb-signup-page-profile-pic-size:88px; --feed-sidebar-padding:32px; --feed-sidebar-width:319px; --feed-polaris-padding-small:32px; --feed-polaris-padding-med:70px; --feed-polaris-padding-large:96px; --feed-polaris-sidebar-width:268px; --feed-width-wide:614px; --feed-with-padding-threshold-min:640px; --font-family-system:-apple-system, BlinkMacSystemFont, &quot;Segoe UI&quot;, Roboto, Helvetica, Arial, sans-serif; --font-weight-system-bold:700; --font-weight-system-extra-bold:800; --font-weight-system-extra-light:200; --font-weight-system-light:300; --font-weight-system-medium:500; --font-weight-system-regular:400; --font-weight-system-semibold:600; --gradient-lavender:#d300c5; --gradient-orange:#ff7a00; --gradient-pink:#ff0169; --gradient-purple:#7638fa; --gradient-yellow:#ffd600; --green-4:#78de45; --green-5:#58c322; --green-6:#37a600; --grey-0:#f5f5f5; --grey-1:#efefef; --grey-2:#dbdbdb; --grey-4:#a8a8a8; --grey-5:#8e8e8e; --grey-6:#737373; --grey-8:#363636; --grey-9:#262626; --grid-column-width-modal:6.25%; --grid-column-width:5.55556%; --hscc-collapse-transition-duration:0.3s; --ig-badge:255, 48, 64; --ig-close-friends-refreshed:28, 209, 79; --ig-disabled-action-text:169, 219, 255; --ig-error-or-destructive:237, 73, 86; --ig-facebook-blue:53, 121, 234; --ig-full-screen-background:54, 54, 54; --ig-live-badge:255, 1, 105; --ig-primary-button:0, 149, 246; --ig-primary-button-hover:24, 119, 242; --ig-secondary-button-background:239, 239, 239; --ig-secondary-button-hover:219, 219, 219; --ig-secondary-button-focused:224, 241, 255; --ig-tertiary-button-background:255, 255, 255; --ig-tertiary-button-border:219, 219, 219; --ig-tertiary-button-hover:245, 245, 245; --ig-tertiary-button-text:38, 38, 38; --ig-subscribers-only:118, 56, 250; --ig-success:88, 195, 34; --ig-text-on-color:255, 255, 255; --ig-text-on-media:255, 255, 255; --igui-border-radius:3px; --in-feed-story-item-height:240px; --in-feed-story-item-width:135px; --input-border-radius:6px; --large-layout-min:1500px; --large-screen-min:876px; --like-animation-duration:1000ms; --live-video-border-radius:4px; --live-video-right-col-width:336px; --media-content-card-width:350px; --media-content-card-width-small:300px; --media-info:335px; --medium-layout-max:1499px; --medium-layout-min:1080px; --medium-screen-max:875px; --medium-screen-min:736px; --mobile-cluster-card-height:74px; --mobile-cluster-card-width:108px; --mobile-grid-item-margin:2px; --mobile-nav-height:45px; --modal-backdrop-dark:rgba(0, 0, 0, 0.85); --modal-backdrop-default:rgba(0, 0, 0, 0.65); --modal-border-radius:12px; --modal-padding:16px; --modal-z-index:100; --nav-narrow-width:72px; --nav-medium-width:244px; --nav-wide-width:335px; --nav-bottom-screen-max:767px; --nav-narrow-screen-min:768px; --nav-medium-screen-min:1264px; --nav-wide-screen-min:1920px; --orange-5:#fd8d32; --photo:600px; --pink-5:#d10869; --post-modal-large-screen-min:1536px; --post-modal-small-screen-max:1366px; --post-step-indicator:168, 168, 168; --purple-5:#a307ba; --red-4:#ff6874; --red-5:#ed4956; --red-6:#c62330; --red-7:#a70311; --reels-large-screen-min:1366px; --refinement-section-height:50px; --revamp-nav-bottom-toolbar-height:50px; --revamp-feed-card-max-height:758px; --revamp-feed-card-min-height:615px; --revamp-feed-card-min-width:710px; --revamp-feed-card-media-min-width:390px; --revamp-feed-card-details-section-width:320px; --revamp-feed-card-details-section-width-xl:340px; --revamp-feed-horizontal-padding-small-screen:24px; --revamp-feed-horizontal-padding-large-screen:32px; --revamp-feed-vertical-padding:32px; --right-rail-width:300px; --scrollable-content-header-height-large:56px; --scrollable-content-header-height-med:49px; --scrollable-content-header-height:44px; --search-box-height:40px; --search-modal-height-expanded:450px; --search-modal-height:362px; --search-modal-top-offset:12px; --search-result-height:50px; --search-result-inline-top-offset:60px; --search-result-list-width:375px; --site-padding-top:30px; --site-width-narrow:600px; --site-width-wide:935px; --small-layout-max:1079px; --small-layout-min:800px; --small-screen-max:735px; --small-screen-min:414px; --story-cube-shading-duration:200ms; --story-desktop-margin-bottom:26px; --story-desktop-margin-top:82px; --story-gallery-preview-scale-correction:2.5; --story-progressbar-update-tick:0.1s; --story-swap-animation-duration:350ms; --system-10-font-size:10px; --system-10-line-height:12px; --system-11-font-size:11px; --system-11-line-height:13px; --system-12-font-size:12px; --system-12-line-height:16px; --system-14-font-size:14px; --system-14-line-height:18px; --system-16-font-size:16px; --system-16-line-height:24px; --system-18-font-size:18px; --system-18-line-height:24px; --system-20-font-size:20px; --system-20-line-height:25px; --system-22-font-size:22px; --system-22-line-height:26px; --system-24-font-size:24px; --system-24-line-height:27px; --system-26-font-size:26px; --system-26-line-height:28px; --system-28-font-size:28px; --system-28-line-height:32px; --system-30-font-size:30px; --system-30-line-height:36px; --system-32-font-size:32px; --system-32-line-height:40px; --web-always-black:0, 0, 0; --web-always-white:255, 255, 255; --web-overlay-on-media:38, 38, 38; --web-secondary-action:224, 241, 255; --yellow-5:#fdcb5c; --challenge-link:54,54,54; --docpen-lightgrey:243,243,243; --ig-banner-background:255,255,255; --ig-elevated-background:255,255,255; --ig-elevated-separator:219,219,219; --ig-focus-stroke:168,168,168; --ig-highlight-background:239,239,239; --ig-link:0,55,107; --ig-primary-background:255,255,255; --ig-primary-text:38,38,38; --ig-secondary-background:250,250,250; --ig-secondary-button:38,38,38; --ig-secondary-text:142, 142, 142; --ig-separator:219,219,219; --ig-stroke:219,219,219; --ig-temporary-highlight:245,251,255; --ig-tertiary-text:199,199,199; --post-separator:239,239,239; --tos-box-shadow:0,0,0;" v-else-if="!this.errormsg">
    <div>
      <div class="ctn-1">
        <div class="ctn-2"></div>
        <div class="ctn-3">
          <div class="ctn-4">
            <div class="ctn-5">
              <div class="ctn-6">
                <div class="ctn-7">
                  <div style="display: flex; flex-direction: column; height: 100%; max-width: 100%;">
                    <div style="max-height: 400px; min-height: 260px;">
                      <div>
                        <div class="ctn-8">
                          <div class="ctn-9" style="height: 100%; width: 100%;">
                            <h1 class="ctn-10" style="width: calc(100% - 100px);">
                              <div style="margin-left:42px">Banned</div>
                            </h1>
                            <div class="ctn-11">
                              <div class="ctn-12">
                                  <button class="ctn-13" type="button" @click="this.$router.go(-1)">
                                      <div class="cth-14">
                                          <svg aria-label="Close" color="#262626" fill="#262626" height="18" role="img" viewBox="0 0 24 24" width="18">
                                              <polyline fill="none" points="20.643 3.357 12 12 3.353 20.647" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="3">
                                              </polyline>
                                              <line fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="3" x1="20.649" x2="3.354" y1="20.649" y2="3.354">
                                              </line>
                                          </svg>
                                      </div>
                                  </button>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                      <div>
                        <div style="max-height: 358px; overflow-y: auto;">
                          <div style="position: relative; display: flex; flex-direction: column; padding-bottom: 0px; padding-top: 0px;" v-for="user in this.banned" :key="user">
                            <div class="ctn-15">
                              <router-link :to="`/users/${user.id}`">
                                <span class="ctn-16">
                                  <div>{{ user.username }}</div>
                                </span>
                              </router-link>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<style>
a {
  text-decoration: none;
}

.ctn-1 {
  display: flex;
  flex-direction: column;
  min-height: 80vh;
}

.ctn-2 {
  bottom: 0;
  right: 0;
  position: fixed;
  left: 0;
  top: 0;
  background-color: var(--overlay-alpha-80);
}

.ctn-3 {
  display: inherit;
  min-height: inherit;
}

.ctn-4 {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  justify-content: center;
}

.ctn-5 {
  display: flex;
  justify-content: center;
}

.ctn-6 {
  position: relative;
  width: 400px;
}

.ctn-7 {
  background-color: rgb(var(--ig-elevated-background));
  border-top-left-radius: var(--modal-border-radius);
  border-top-right-radius: var(--modal-border-radius);
  border-bottom-right-radius: var(--modal-border-radius);
  border-bottom-left-radius: var(--modal-border-radius);
}

.ctn-8 {
  align-items: center;
  border-bottom: 1px solid rgb(var(--ig-elevated-separator));
  box-sizing: border-box;
  display: inline-block;
  height: calc( var(--scrollable-content-header-height) - 1px );
  position: relative;
  width: 100%;
}

.ctn-9 {
  display: flex;
  align-items: center;
}

.ctn-10 {
  align-items: center;
  border: 0;
  color: rgb(var(--ig-primary-text));
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  font: inherit;
  font-size: var(--system-16-font-size);
  font-weight: var(--font-weight-system-semibold);
  justify-content: center;
  line-height: var(--system-16-line-height);
  margin: 0;
  min-width: 0;
  padding: 0;
  text-align: center;
  vertical-align: baseline;
}

.ctn-11 {
  align-items: center;
  display: flex;
  flex-basis: 48px;
  flex-direction: column;
  height: 100%;
  justify-content: center;
  position: relative;
  float: right;
}

.ctn-12 {
  display: flex;
  align-content: stretch;
  align-items: stretch;
  flex-direction: column;
  justify-content: flex-start;
  flex: 0 0 auto;
  padding-left: calc(var(--base-unit) * 2);
  padding-right: calc(var(--base-unit) * 2);
  box-sizing: border-box;
  position: relative;
}

.ctn-13 {
  align-items: center;
  background: transparent;
  border: none;
  cursor: pointer;
  display: flex;
  justify-content: center;
  padding: 8px;
}

.ctn-14 {
  align-items: center;
  display: flex;
  justify-content: center;
}

.ctn-15 {
  padding-left: calc(var(--base-unit) * 4);
  padding-right: calc(var(--base-unit) * 4);
  padding-bottom: calc(var(--base-unit) * 2);
  padding-top: calc(var(--base-unit) * 2);
}

.ctn-16 {
  font-family: var(--font-family-system);
  font-size: var(--system-14-font-size);
  line-height: var(--system-14-line-height);
  margin: -3px 0 -4px;
  font-weight: var(--font-weight-system-semibold);
  color: rgb(var(--ig-primary-text));
  display: inline!important;
  margin: 0!important;
}
</style>
