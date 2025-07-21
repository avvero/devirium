package pw.avvero;

import com.badlogic.gdx.ApplicationAdapter;
import com.badlogic.gdx.Gdx;
import com.badlogic.gdx.Input.Keys;
import com.badlogic.gdx.graphics.GL20;
import com.badlogic.gdx.graphics.OrthographicCamera;
import com.badlogic.gdx.graphics.Color;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer;
import com.badlogic.gdx.graphics.glutils.ShapeRenderer.ShapeType;
import com.badlogic.gdx.math.Rectangle;
import com.badlogic.gdx.math.Vector2;
import com.badlogic.gdx.utils.Array;
import com.badlogic.gdx.utils.ScreenUtils;
import com.badlogic.gdx.physics.box2d.*;

/** {@link com.badlogic.gdx.ApplicationListener} implementation shared by all platforms. */
public class Main extends ApplicationAdapter {
    private static final float GRAVITY = -9.8f;
    private static final float PLAYER_SPEED = 5f;
    private static final float JUMP_VELOCITY = 5f;
    private static final float SCALE = 100f; // Box2D works in meters
    
    private OrthographicCamera camera;
    private ShapeRenderer shapeRenderer;
    
    private World world;
    private Box2DDebugRenderer debugRenderer;
    
    private Body playerBody;
    private Array<Body> platforms = new Array<Body>();
    
    private boolean canJump = false;
    
    @Override
    public void create() {
        // Create camera
        camera = new OrthographicCamera();
        camera.setToOrtho(false, 800, 480);
        
        // Create shape renderer
        shapeRenderer = new ShapeRenderer();
        
        // Create Box2D world with gravity
        world = new World(new Vector2(0, GRAVITY), true);
        debugRenderer = new Box2DDebugRenderer();
        
        // Add contact listener to detect when player is on ground
        world.setContactListener(new ContactListener() {
            @Override
            public void beginContact(Contact contact) {
                Fixture fixtureA = contact.getFixtureA();
                Fixture fixtureB = contact.getFixtureB();
                
                if (fixtureA.getUserData() != null && fixtureA.getUserData().equals("player") ||
                    fixtureB.getUserData() != null && fixtureB.getUserData().equals("player")) {
                    canJump = true;
                }
            }
            
            @Override
            public void endContact(Contact contact) {}
            
            @Override
            public void preSolve(Contact contact, Manifold oldManifold) {}
            
            @Override
            public void postSolve(Contact contact, ContactImpulse impulse) {}
        });
        
        // Create player
        createPlayer();
        
        // Create platforms
        createPlatforms();
    }
    
    private void createPlayer() {
        // Create player body definition
        BodyDef bodyDef = new BodyDef();
        bodyDef.type = BodyDef.BodyType.DynamicBody;
        bodyDef.position.set(1.5f, 3f);
        
        // Create player body
        playerBody = world.createBody(bodyDef);
        
        // Create player shape
        PolygonShape shape = new PolygonShape();
        shape.setAsBox(0.2f, 0.2f);
        
        // Create player fixture
        FixtureDef fixtureDef = new FixtureDef();
        fixtureDef.shape = shape;
        fixtureDef.density = 1.0f;
        fixtureDef.friction = 0.4f;
        fixtureDef.restitution = 0.1f;
        
        Fixture fixture = playerBody.createFixture(fixtureDef);
        fixture.setUserData("player");
        
        // Dispose shape
        shape.dispose();
    }
    
    private void createPlatforms() {
        // Create ground platform
        createPlatform(4f, 0.5f, 8f, 0.5f);
        
        // Create some platforms
        createPlatform(1f, 2f, 2f, 0.25f);
        createPlatform(5f, 2.5f, 2f, 0.25f);
        createPlatform(3f, 4f, 1.5f, 0.25f);
        
        // Create world boundaries
        // Left wall
        createPlatform(-0.5f, 2.4f, 1f, 5f);
        // Right wall
        createPlatform(8.5f, 2.4f, 1f, 5f);
        // Ceiling (to prevent going too high)
        createPlatform(4f, 6f, 10f, 1f);
    }
    
    private void createPlatform(float x, float y, float width, float height) {
        // Create platform body definition
        BodyDef bodyDef = new BodyDef();
        bodyDef.type = BodyDef.BodyType.StaticBody;
        bodyDef.position.set(x, y);
        
        // Create platform body
        Body body = world.createBody(bodyDef);
        
        // Create platform shape
        PolygonShape shape = new PolygonShape();
        shape.setAsBox(width / 2, height / 2);
        
        // Create platform fixture
        FixtureDef fixtureDef = new FixtureDef();
        fixtureDef.shape = shape;
        fixtureDef.friction = 0.5f;
        
        body.createFixture(fixtureDef);
        platforms.add(body);
        
        // Dispose shape
        shape.dispose();
    }

    @Override
    public void render() {
        // Clear screen
        ScreenUtils.clear(0.4f, 0.6f, 1f, 1f);
        
        // Update physics world
        float deltaTime = Gdx.graphics.getDeltaTime();
        world.step(1/60f, 6, 2);
        
        // Process input
        processInput();
        
        // Update camera
        camera.update();
        
        // Set projection matrix for shape renderer
        shapeRenderer.setProjectionMatrix(camera.combined);
        
        // Draw player
        Vector2 playerPos = playerBody.getPosition();
        shapeRenderer.begin(ShapeType.Filled);
        shapeRenderer.setColor(Color.RED);
        shapeRenderer.rect(
            (playerPos.x - 0.2f) * SCALE, 
            (playerPos.y - 0.2f) * SCALE, 
            0.4f * SCALE, 
            0.4f * SCALE
        );
        
        // Draw platforms
        shapeRenderer.setColor(Color.BROWN);
        for (Body platform : platforms) {
            Vector2 pos = platform.getPosition();
            // Get the width and height from the fixture
            PolygonShape shape = (PolygonShape) platform.getFixtureList().get(0).getShape();
            Vector2 size = new Vector2();
            shape.getVertex(1, size);
            size.x *= 2; // Convert half-width to width
            size.y *= 2; // Convert half-height to height
            
            shapeRenderer.rect(
                (pos.x - size.x/2) * SCALE, 
                (pos.y - size.y/2) * SCALE, 
                size.x * SCALE, 
                size.y * SCALE
            );
        }
        shapeRenderer.end();
        
        // Debug rendering
        // debugRenderer.render(world, camera.combined.scl(SCALE));
    }
    
    private void processInput() {
        // Get current velocity
        Vector2 vel = playerBody.getLinearVelocity();
        Vector2 pos = playerBody.getPosition();
        
        // Check if player is out of bounds and reset position if needed
        if (pos.x < 0 || pos.x > 8 || pos.y < 0 || pos.y > 6) {
            // Reset to a safe position
            playerBody.setTransform(1.5f, 3f, 0);
            playerBody.setLinearVelocity(0, 0);
            return;
        }
        
        // Reset horizontal velocity
        playerBody.setLinearVelocity(0, vel.y);
        
        // Apply left/right movement
        if (Gdx.input.isKeyPressed(Keys.LEFT)) {
            playerBody.setLinearVelocity(-PLAYER_SPEED, vel.y);
        }
        if (Gdx.input.isKeyPressed(Keys.RIGHT)) {
            playerBody.setLinearVelocity(PLAYER_SPEED, vel.y);
        }
        
        // Apply jump if on ground
        if (Gdx.input.isKeyJustPressed(Keys.SPACE) && canJump) {
            playerBody.setLinearVelocity(vel.x, JUMP_VELOCITY);
            canJump = false;
        }
    }

    @Override
    public void dispose() {
        shapeRenderer.dispose();
        world.dispose();
    }
}
