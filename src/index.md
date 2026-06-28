---
layout: layouts/base.njk
title: Home
---

<section>
	<header>
		<h1>Zachary Norman</h1>
		<p>Developer • frontend & backend • security-minded</p>
	</header>

  <section aria-labelledby="posts">
    <h2 id="posts">Latest Posts</h2>
    <ul>
      {% for post in collections.posts %}
      <li><a href="{{ post.url }}">{{ post.data.title }}</a></li>
      {% endfor %}
    </ul>
  </section>
			<p>I study Computer Science at the University of Guelph. I build full-stack apps, machine learning prototypes, and secure services. I prefer small, maintainable code and measurable results.</p>
		</div>
	</section>

	<section aria-labelledby="services">
		<h2 id="services">Services</h2>
		<ul>
			<li><strong>Full Stack Development</strong> — MERN / PERN, production-ready apps.</li>
			<li><strong>Algorithm Engineering</strong> — performant implementations in Python/Java/C#.</li>
			<li><strong>Machine Learning</strong> — prototypes with TensorFlow / PyTorch.</li>
			<li><strong>Ethical Hacking</strong> — secure by design, pen-test experience.</li>
		</ul>
	</section>

	<section aria-labelledby="portfolio">
		<h2 id="portfolio">Portfolio</h2>
		<div class="grid">
			<figure>
				<img src="/assets/images/rbc.png" alt="Full Stack WebApp">
				<figcaption>Full Stack WebApp Development</figcaption>
			</figure>
			<figure>
				<img src="/assets/images/ai-icon.png" alt="Facial Detection">
				<figcaption>Facial Detection (ML)</figcaption>
			</figure>
			<figure>
				<img src="/assets/images/java-icon.png" alt="Algorithm Engineering">
				<figcaption>Algorithm Engineering</figcaption>
			</figure>
		</div>
	</section>

	<section aria-labelledby="contact">
		<h2 id="contact">Contact</h2>
		<p>Email: znorman(at)uoguelph.ca</p>
		{% include "partials/contact-form.njk" %}
	</section>
</section>
