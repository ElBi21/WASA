<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {}
</script>

<template>
	<header class="navbar sticky-top flex-md-nowrap p-0">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-5" href="#/">WASAText</a>
		<div class="navbar-options">
            <img class="navbar-icon" src="../public/icons/gear-solid.svg" alt="Settings icon">
            <div class="navbar-user-icon"></div>
        </div>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">

                    <div class="chat-preview">
                        <img class="chat-preview-pfp" src="../public/test_assets/group_pfp.png" alt="Chat pfp">
                        <div class="chat-preview-data">
                            <p class="chat-preview-title">Amici del bunker</p>
                            <p class="chat-preview-last-msg"><b>Author</b>: This is the last message</p>
                        </div>
                    </div>

                    <div class="chat-preview">
                        <img class="chat-preview-pfp" src="../public/test_assets/group_pfp.png" alt="Chat pfp">
                        <div class="chat-preview-data">
                            <p class="chat-preview-title">Amici del bunker</p>
                            <p class="chat-preview-last-msg"><b>Author</b>: This is the last message</p>
                        </div>
                    </div>

                    <div class="chat-preview">
                        <img class="chat-preview-pfp" src="../public/test_assets/group_pfp.png" alt="Chat pfp">
                        <div class="chat-preview-data">
                            <p class="chat-preview-title">Amici del bunker</p>
                            <p class="chat-preview-last-msg"><b>Author</b>: This is the last message</p>
                        </div>
                    </div>



                    <!-- <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								Home
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/link1" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
								Menu item 1
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/link2" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
								Menu item 2
							</RouterLink>
						</li>
					</ul>

					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Secondary menu</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink :to="'/some/' + 'variable_here' + '/path'" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
								Item 1
							</RouterLink>
						</li>
					</ul>-->
				</div>
			</nav>

			<main class="main-view col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style>
</style>
