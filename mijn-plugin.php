<?php

require_once(plugin_dir_path(__FILE__) . '/lib/autoload.php');

use ElenbaasIt\MijnPlugin\MijnPlugin;
use ElenbaasIt\MijnPlugin\MijnPlugin_Activator;
use ElenbaasIt\MijnPlugin\MijnPlugin_Deactivator;

/**
 * @wordpress-plugin
 * Plugin Name:       Mijn Plugin
 * Description:       A WordPress plugin
 * Version:           0.1.0
 * Author:            Your Name 
 * Text Domain:       mijn-plugin
 * Domain Path:       /languages
 * Requires PHP:      8.1.0
 */

// If this file is called directly, abort.
if (! defined('WPINC')) {
    die;
}

define('MIJN_PLUGIN_VERSION', '0.10.0');

function activate_mijn_plugin() {
    MijnPlugin_Activator::activate();
}

function deactivate_mijn_plugin() {
    MijnPlugin_Deactivator::deactivate();
}

register_activation_hook(__FILE__, 'activate_mijn_plugin');
register_deactivation_hook(__FILE__, 'deactivate_mijn_plugin');

/**
 * Begins execution of the plugin.
 *
 * Since everything within the plugin is registered via hooks,
 * then kicking off the plugin from this point in the file does
 * not affect the page life cycle.
 *
 * @since    1.0.0
 */
function run_mijn_plugin() {

    $plugin = new MijnPlugin();
    $plugin->run();
}
run_mijn_plugin();
